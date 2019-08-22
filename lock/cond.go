package lock

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
sync.Cond类型即是Go中的条件变量，该类型内部包含一个锁接口，以此来实现锁操作。

条件变量必须使用下列函数才能创建：

func NewCond(l locker) *Cond        // 条件变量必须传入一个锁，二者需要配合使用
*sync.Cond类型有三个方法：

Wait: 等待通知。对锁进行解锁，并且使所在的协程阻塞，一旦收到通知，则唤醒，并立即锁定该所
Signal: 发送通知(单发)
Broadcast: 发送通知(广播）
*/

func condMain() {
	cond := sync.NewCond(&sync.RWMutex{})
	buf := make([]byte , 0)

	for i := 0; i < 5; i++ {
		go func(a int) {
			time.Sleep(time.Second * time.Duration(2))
			log.Println(a , " got the lock")
			cond.L.Lock()
			buf = append(buf , []byte(fmt.Sprintf("%d" , a))...)
			cond.Broadcast()
			cond.L.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		go func(a int) {
			cond.L.Lock()
			log.Println("R -" , a , " ready")
			data := string(buf)
			buf = make([]byte , 0)
			log.Println("R -" , a ," got " , data)
			cond.Wait()
		}(i)
	}

	for {}

}