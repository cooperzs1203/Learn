package conditionStatement

import (
	"log"
	"time"
)

func conditionStateMent() {
	ifStateMent()
	switchStateMent()
	selectStateMent()
}

func ifStateMent() {
	a , b := 1 ,2
	if a > b {
		log.Printf("%d > %d" , a , b)
	}else if a == b {
		log.Printf("%d == %d" , a , b)
	}else {
		log.Printf("%d < %d" , a , b)
	}
}

func switchStateMent() {
	a := 5
	switch a {
	case 1:
		log.Println("1")
	case 2:
		log.Println("2")
	case 3:
		log.Println("3")
	case 4:
		log.Println("4")
	case 5:
		log.Println("5")
	default:
		log.Println("Unknow Number")
	}
}

func selectStateMent() {
	intChan := make(chan int , 0)
	go func() {
		time.Sleep(time.Second * time.Duration(3))
		intChan <- 1
	}()

	for {
		select {
		case a := <-intChan:
			log.Println(a)
			goto BreakPoint
		case <-time.After(time.Duration(1) * time.Second):
			log.Println("wait")
		}
	}

	BreakPoint:
		log.Println("Finish")
}