package lock

import (
	"log"
	"sync"
)

var (
	once sync.Once
	theStu *Student
)

type Student struct {
	Name string
	Age int64
	Sex bool
}

func TheStudent() *Student {
	once.Do(func() {
		theStu = &Student{"cooper" , 23 , true}
	})

	return theStu
}

func onceMain() {
	astu := TheStudent()
	bstu := TheStudent()
	cstu := TheStudent()

	log.Printf("%p" , astu)
	log.Printf("%p" , bstu)
	log.Printf("%p" , cstu)
}