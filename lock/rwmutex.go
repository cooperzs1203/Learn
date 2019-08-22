package lock

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
在开发场景中，经常遇到多处并发读取，一次并发写入的情况，Go为了方便这些操作，在互斥锁基础上，提供了读写锁操作。

读写锁即针对读写操作的互斥锁，简单来说，就是将数据设定为 写模式（只写）或者读模式（只读）。使用读写锁可以分别针对读操作和写操作进行锁定和解锁操作。

读写锁的访问控制规则与互斥锁有所不同：
1.写操作与读操作之间也是互斥的
2.读写锁控制下的多个写操作之间是互斥的，即一路写
3.多个读操作之间不存在互斥关系，即多路读

在Go中，读写锁由结构体sync.RWMutex表示，包含两对方法：

// 设定为写模式：与互斥锁使用方式一致，一路只写
func (*RWMutex) Lock()				// 锁定写
func (*RWMutex) Unlock()			// 解锁写

// 设定为读模式：对读执行加锁解锁，即多路只读
func (*RWMutex) RLock()
func (*RWMutex) RUnlock()
*/
func rwMutexMain() {
	//rwfirst()
	rwSecond()
}

type SafeMap struct {
	Conns map[string]string
	rwMutex *sync.RWMutex
}

func (sm *SafeMap) Set(key, value string) {
	sm.rwMutex.Lock()
	defer sm.rwMutex.Unlock()
	sm.Conns[key] = value
}

func (sm *SafeMap) Get(key string) (value string , ok bool) {
	sm.rwMutex.RLock()
	defer sm.rwMutex.RUnlock()
	value , ok = sm.Conns[key]
	return
}

func rwfirst() {
	var rwMutex sync.RWMutex

	for i := 0; i < 3; i++ {
		go func(a int) {
			rwMutex.RLock()
			log.Println(a)
			time.Sleep(time.Second * time.Duration(2))
			rwMutex.RUnlock()
		}(i)
	}

	time.Sleep(time.Second * time.Duration(10))
	rwMutex.Lock()
}

func rwSecond() {
	sm := SafeMap{
		make(map[string]string),
		&sync.RWMutex{},
	}

	for i := 0; i < 10; i++ {
		go func(a int) {
			time.Sleep(time.Second * time.Duration(2))
			v , ok := sm.Get(fmt.Sprintf("%d" , a))
			log.Println(a , "--" , v , ok)
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(a int) {
			sm.Set(fmt.Sprintf("%d",a) , fmt.Sprintf("%d" , a))
			log.Println("Set --" , a)
		}(i)
	}

	for {}

}