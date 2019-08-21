package function

import (
	"fmt"
	"log"
)

/*
func function_name( [parameter list] ) [return_types] {
   函数体
}
*/

func functionMain() {
	log.Println(normalFunc(5 , 3))

	a , b := 5 , 3
	pointFunc(&a , &b)
	log.Println(a , b)
	indefiniteParam("cooper" , 5 , 0 , 4 , 5)

	f := func(a int) int { return a +2 }
	funcParam(f)

	c := blockFunc(5)
	log.Println(c())
	log.Println(c())

	var cI cInt
	cI = 9
	log.Println(cI.addMethod(6))

}

func normalFunc(a, b int) (int, int) {
	x := a + b
	y := a - b
	return x , y
}

func pointFunc(a, b *int) {
	*a = *a + *b
	*b = *a - *b
}

func indefiniteParam(str string, a ...int) {
	s := ""
	s += str
	for _ , v := range a {
		s += fmt.Sprintf("%d" , v)
	}
	log.Println(s)
}

func funcParam(f func(int) int) {
	log.Println(f(8))
}

func blockFunc(n int) func() int {
	num := n
	return func() int {
		num++
		return num
	}
}

type cInt int

func (ci cInt) addMethod(n int) int {
	return int(ci) + n
}