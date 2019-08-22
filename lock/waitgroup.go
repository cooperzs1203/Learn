package lock

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
sync.WaitGroup类型的值也是并发安全的，该类型结构体中内内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少 。

当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减1。 同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时， 表示所有任务己经完成。

等待组常用方法：

(wg *WaitGroup) Add(delta int)	等待组计数器+1，该方法也可以传入负值让等待计数减少，切记不能减少为负数，会引发崩溃
(wg *WaitGroup) Done()	等待组计数器-1，等同于Add传入负值
(wg *WaitGroup) Wait()	等待组计数器!=0时阻塞，直到为0
应用场景：WaitGroup一般用于协调多个goroutine运行。
*/

func waitGroup() {
	//wgFirst()
	wgSecond()
}

func wgFirst() {
	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a int) {
			t := rand.Intn(10)
			time.Sleep(time.Second * time.Duration(t))
			log.Println(a)
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("we done")
}

func wgSecond() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	money := 0

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a int) {
			mutex.Lock()
			defer mutex.Unlock()
			log.Println(a , "Got Lock")
			for j := 0; j < 100; j++ {
				money += 10
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println(money)
}
