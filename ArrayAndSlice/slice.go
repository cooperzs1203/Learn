package ArrayAndSlice

import "log"

/*
切片("动态数组"),与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
var identifier []type

也可以指定容量，其中capacity为可选参数。
make([]T, length, capacity)

*/

var intSlice []int = make([]int , 0 , 5)

func sliceMain() {
	log.Println(intSlice)
	for i := 0; i<cap(intSlice); i++ {
		intSlice = append(intSlice , i * 100)
	}
	log.Println(intSlice)
	log.Println(&intSlice[0])

	// [startIndex:endIndex] =>[startIndex:endIndex)
    fIntSlice := intSlice[:]
	log.Println(fIntSlice)
	log.Println(&fIntSlice[0])
    log.Println(fIntSlice[1:3])

    log.Println(len(fIntSlice))
    log.Println(cap(fIntSlice))

    tailSlice(fIntSlice)
	tailSlicePoint(&fIntSlice)
    nilSlice()
    copySlice()
    appendSlice()
}

func tailSlice(intSlice []int) {
	log.Println(intSlice)
	log.Println(&intSlice[0])
}

func tailSlicePoint(intSlice *[]int) {
	log.Println(*intSlice)
	log.Println(&(*intSlice)[0])
}

func nilSlice() {
	var nilSlice []string
	log.Println(nilSlice)
	log.Println(nilSlice==nil)
}

// copy 的dst len 不能为空
func copySlice() {
	aSlice := []int{1,2,3,4,5,6,7,8,9,0}
	log.Println(aSlice)
	log.Println(&aSlice[0])

	bSlice := make([]int , len(aSlice) , cap(aSlice))
	copy(bSlice , aSlice)
	log.Println(bSlice)
	log.Println(&bSlice[0])
}

func appendSlice() {
	a := []int{1,2,3,4,5}
	a = append(a , 6)
	log.Println(a)
	a = append(a , 7,8,9)
	log.Println(a)
	b := []int{10,11,12}
	a = append(a , b...)
	log.Println(a)
}