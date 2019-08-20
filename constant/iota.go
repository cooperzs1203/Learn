package constant

import "log"

const (
	A = iota
	B
	C
)

const (
	D = iota
	E
	F
	G = "G"
	H
	I = 100
	J = iota
	K
)

// << 左移操作，将数字转为二进制，左移，右补0
const (
	left = 1<<iota
	right = 3<<iota
	up
	down
)

func iotaFunc() {
	log.Println(A)
	log.Println(B)
	log.Println(C)
	log.Println(D)
	log.Println(E)
	log.Println(F)
	log.Println(G)
	log.Println(H)
	log.Println(I)
	log.Println(J)
	log.Println(K)
	log.Println(left)
	log.Println(right)
	log.Println(up)
	log.Println(down)
}