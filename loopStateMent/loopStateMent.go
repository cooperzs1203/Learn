package loopStateMent

import "fmt"

func loopStateMent() {
	forStateMent()
	loopControl()
	foreverLoop()
}

func forStateMent() {
	kv := make(map[string][]int)
	for i := 0; i < 10; i++ {
		list := make([]int , 0)
		for j := 10; j < 20; j++ {
			list = append(list , j+1)
		}
		kv[fmt.Sprintf("%d" , i+1)] = list
	}

	for key , value := range kv {
		for index , value := range value {
			fmt.Println(key , " - " , index , " - ", value)
		}
	}
}

func loopControl() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 20; i++ {
		if i == 15 { break }
		fmt.Println(i)
	}

	for i := 0; i < 20; i++ {
		if i == 5 {
			goto End
		}
		fmt.Println(i)
	}

	End:
}

func foreverLoop() {
	for {
		fmt.Println("forever")
	}
}