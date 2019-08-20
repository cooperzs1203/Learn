package constant

import "log"

// const identifier [type] = value
const s string = "constant"
const i , j = 5 , 6

const (
	a = "cooper"
)

func constant() {
	log.Println(s)
	log.Println(a)
	log.Printf("%d : %d" , i , j)
}