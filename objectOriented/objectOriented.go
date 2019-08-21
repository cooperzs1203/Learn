package objectOriented

import (
	"./object"
	"log"
)

func objectOrientedMain() {
	box()
	inherit()
}

func box() {
	stu := object.Student("cooper")
	log.Println(stu)
	stu.SetId("123456789")
	log.Println(stu.GetId())
	log.Println(stu)
}

type Person struct {
	Name string
	Age int
}

func (p *Person) Introduce() {
	log.Printf("I'm %v , %d years old." , p.Name , p.Age)
}

type Male struct {
	Person
	DickLong int
}

func inherit() {
	guy := Male{}
	guy.Name = "cooper"
	guy.Age = 26
	guy.DickLong = 18
	guy.Introduce()
}