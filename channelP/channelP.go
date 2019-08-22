package channelP

import (
	"log"
	"time"
)

func channelPMain() {
	//channelA()
	//channelB()
	//channelC()
	channelD()
}

func channelA() {
	achan := make(chan int)

	go func() {
		time.Sleep(time.Second*time.Duration(2))
		achan <- 0
		log.Println("continue")
	}()

	data , ok := <- achan
	log.Println(ok , data)
	log.Println("Finish")
}

func channelB() {
	achan := make(chan int)

	go func() {
		time.Sleep(time.Second*time.Duration(2))
		achan <- 0
		log.Println("continue")
	}()

	data := <- achan
	log.Println(data)
	log.Println("Finish")
}

func channelC() {
	achan := make(chan int , 3)

	achan <- 0
	achan <- 1

	log.Println(len(achan))
	log.Println(cap(achan))
}

func channelD() {
	achan := make(chan int)

	go func(ch <-chan int) {
		for {
			d := <- ch
			log.Println("go func " , d)
		}
	}(achan)

	go func(ch chan <- int) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * time.Duration(1))
			ch <- i
		}
	}(achan)

	for {}
}