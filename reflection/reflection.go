package reflection

import (
	"fmt"
	"log"
	"reflect"
)

type Student struct {
	Name 	string 	`json:"name" form:"name"`
	Age  	int    	`json:"age" form:"age"`
	Sex  	bool    `json:"sex" form:"sex"`
	Grade 	int   	`json:"grade" form:"grade"`
	Class 	int   	`json:"class" form:"class"`
	Id    	string 	`json:"id" form:"id"`
}

func (s *Student) Info() string {
	return fmt.Sprintf("%v : %d-%d : %v" , s.Id , s.Grade , s.Class , s.Name)
}

func reflectionMain() {
	stu := &Student{
		"jack",
		18,
		true,
		11,
		3,
		"55648521355569",
	}
	//typeOf(stu)
	//valueOf(stu)
	//setValue(stu)
	//log.Println(stu)
	callMethod(stu)
}

func typeOf(d interface{}) {
	t := reflect.TypeOf(d)
	t = t.Elem() // 获取原始值对应的反射对象
	log.Println(t.Kind() , t.Name())

	for i := 0; i < t.NumField(); i++ {
		log.Println(t.Field(i).Name , " -- " , t.Field(i).Tag.Get("form"))
	}
}

func valueOf(d interface{}) {
	v := reflect.ValueOf(d)
	v = v.Elem()
	log.Println(v.Kind())

	for i := 0; i < v.NumField(); i++ {
		log.Println(v.Field(i))
	}
}

func setValue(d interface{}) {
	stu , ok := d.(*Student) // 必须传指针
	if !ok {
		log.Println("exit")
		return
	}

	v := reflect.ValueOf(&stu.Age)
	v = v.Elem()
	log.Println(v.CanSet())
	log.Println(stu)
	v.SetInt(19)
}

func callMethod(d interface{}) {
	v := reflect.ValueOf(d)
	method := v.MethodByName("Info")
	ret := method.Call(make([]reflect.Value , 0))
	log.Println(ret)
}