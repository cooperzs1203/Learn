package operator

import "log"

func operatorMain() {
	mathOperator()
	relationOperator()
	logicOperator()
	byteOperator()
	assignmentOperator()
	anotherOperator()
}

func mathOperator() {
	a , b := 10 , 3
	log.Println(a+b)
	log.Println(a-b)
	log.Println(a*b)
	log.Println(a/b)
	log.Println(a%b)
	a++
	log.Println(a)
	a--
	log.Println(a)
}

func relationOperator() {
	a , b := 5 , 6

	log.Println(a==b)
	log.Println(a!=b)
	log.Println(a>b)
	log.Println(a<b)
	log.Println(a>=b)
	log.Println(a<=b)
}

func logicOperator() {
	a , b := true , false

	log.Println(a&&b)
	log.Println(a||b)
	log.Println(!b)
}

func byteOperator() {
	a , b := 60 , 13
	// a 0011 1100
	// b 0000 1101
	log.Println(a&b) // 0000 1100 => 12
	log.Println(a|b) // 0011 1101 => 61
	log.Println(a^b) // 0011 0001 => 49
	log.Println(a<<2)// 1111 0000 => 240
	log.Println(a>>2)// 0000 1111 => 15
}

func assignmentOperator() {
	var a int
	a = 10
	log.Println(a)
	a += 1
	log.Println(a)
	a -= 1
	log.Println(a)
	a *= 1
	log.Println(a)
	a /= 1
	log.Println(a)
	a %= 2
	log.Println(a)
	a = 10 // 0000 1010
	a <<= 2
	log.Println(a) // 0010 1000 => 40
	a >>= 2
	log.Println(a)
	a &= 2
	log.Println(a) // 0000 0010 => 2
	a |= 2
	log.Println(a) // 0000 0010 => 2
	a ^= 2
	log.Println(a) // 0000 0000 => 0
}

func anotherOperator() {
	b := 55
	log.Println(&b)
	c := &b
	log.Println(*c)
}