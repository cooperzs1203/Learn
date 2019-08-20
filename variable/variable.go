package variable

import "log"

/*
var identifier type
var identifier1, identifier2 type
*/
var (
	a , b = 1 , 2
	s = "cooper"
)

/*
指定变量类型，如果没有初始化，则变量默认为零值
数值类型（包括complex64/128）为 0
布尔类型为 false
字符串为 ""（空字符串）
以下几种类型为 nil：
var a *int
var a []int
var a map[string] int
var a chan int
var a func(string) int
var a error // error 是接口

（notice:这种因式分解关键字的写法一般用于声明全局变量）
*/
var (
	intPoint *int
	intSlice []int
	siMap map[string] int
	chanInt chan int
	funcSi func(string) int
	err error // error 是接口
)

func variableMain() {
	log.Println(a , b , s)

	zeroValue()
	judgeTypeSelf()
	varLess()
}

func zeroValue() {
	var c int
	log.Println(c)
	var d float64
	log.Println(d)
	log.Println(intPoint)
	log.Println(intSlice)
	log.Println(siMap)
	log.Println(chanInt)
	log.Println(funcSi)
	log.Println(err)

	log.Println(intSlice==nil)
	log.Println(siMap==nil)
}

func judgeTypeSelf() {
	var s = 56.4
	log.Println(s)
}

/*
:= 左侧如果没有声明新的变量，就产生编译错误
只能被用在函数体内
*/
func varLess() {
	vl := "abcc"
	log.Println(vl)
	vl1 , vl2 := "xcxc" , 44
	log.Println(vl1)
	log.Println(vl2)
	vl1 , vl3 := "dddd" , 88.98
	log.Println(vl1)
	log.Println(vl3)
}