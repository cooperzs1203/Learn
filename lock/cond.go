package lock

import (
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
	condition := false

	go func() {
		cond.L.Lock()

		for condition == false {
			log.Println("loop")
			cond.Wait()
			log.Println("goroutine")
		}

		cond.L.Unlock()
	}()

	time.Sleep(time.Second * time.Duration(3))
	cond.L.Lock()
	condition = true
	cond.Signal()
	cond.L.Unlock()

	for {}

}