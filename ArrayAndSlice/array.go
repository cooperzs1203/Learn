package ArrayAndSlice

import "log"

/*
数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。
var variable_name [SIZE] variable_type
初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
*/

var floats = [5]float64{1.1 , 2.2 , 3.3 , 4.4 , 5.5}
var ints = [...]int{1,2,3,4,5}

func arrayMain() {
	log.Println(floats)
	log.Println(ints)
	log.Println(ints[4])

	for index , value := range ints {
		log.Printf("%d : %v" , index , value)
	}

	bools := [3]bool{true , true , false}
	log.Println(bools)
	log.Println(&bools[0])
	tailArray(bools)
	tailArrayPoint(&bools)
}

func tailArray(bools [3]bool) {
	log.Println("----")
	log.Println(bools)
	log.Println(&bools[0])
}

func tailArrayPoint(bools *[3]bool) {
	log.Println("----")
	log.Println(bools)
	log.Println(&bools[0])
}
