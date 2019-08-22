package lock

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
互斥锁是传统并发程序进行共享资源访问控制的主要方法。

Go中由标准库sync中的结构体Mutex表示互斥锁，保证同时只有一个 goroutine 可以访问共享资源，该结构体包含两个方法：

Lock：锁定当前互斥量
Unlock：解锁当前互斥量
*/
func mutexMain() {
	//first()
	second()
}

func first() {
	var mutex sync.Mutex
	log.Println("lock")
	mutex.Lock() // 1.此处加锁

	for i := 0; i < 3; i++ {
		go func(a int) {
			log.Println("Now " , a) // 3.各个协程到了争抢锁，堵塞
			mutex.Lock() // 5.3个协程中的某个抢到了锁，其它两个依旧堵塞
			log.Println("loop :" , a)
			//mutex.Unlock() // 6.解锁后，其它两个协程抢锁，然后解锁，剩下的获得锁，然后解锁
		}(i)
	}

	time.Sleep(time.Second) // 2.等各个协程进入抢夺锁状态
	mutex.Unlock() // 4.主线程解锁
	log.Println("unlock main")
	time.Sleep(time.Second)
}

type Account struct {
	Money float64
	mutex *sync.Mutex
}

func (a *Account) save(money float64) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	a.Money += money
}

func (a *Account) query() string {
	return fmt.Sprintf("%.2f" , a.Money)
}

func second() {
	myAccount := Account{0 , &sync.Mutex{}}

	for i := 0; i < 5; i++ {
		go func(money int) {
			log.Println("Save " , money)
			myAccount.save(float64(money))
		}(i)
	}

	for {
		time.Sleep(time.Second)
		log.Println(myAccount.query())
	}

}