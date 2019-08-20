package variable

import "log"

/*
1.所有像 int、float、bool 和 string 这些基本类型都属于值类型，
使用这些类型的变量直接指向存在内存中的值.
2.当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，
实际上是在内存中将 i 的值进行了拷贝.

*/
func referenceMain() {
	i , f , b , s := 5 , 4.5 , true , "cooper"
	log.Println(i , f , b , s)
	log.Println(&i , &f , &b , &s)

	log.Println("--------------")
	i2 , f2 , b2 , s2 := i , f , b , s
	log.Println(i2 , f2 , b2 , s2)
	log.Println(&i2 , &f2 , &b2 , &s2)

	log.Println("--------------")
	i3 , f3 , b3 , s3 := &i , &f , &b , &s
	log.Println(*i3 , *f3 , *b3 , *s3)
	log.Println(i3 , f3 , b3 , s3)

	log.Println("--------------")
	*i3 , *f3 , *b3 , *s3 = 6 , 6.5 , false , "ddd"
	log.Println(i , f , b , s)
	log.Println(&i , &f , &b , &s)
	log.Println(*i3 , *f3 , *b3 , *s3)
	log.Println(i3 , f3 , b3 , s3)
}
