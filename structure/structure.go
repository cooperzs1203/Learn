package structure

import (
	"fmt"
	"log"
)

func structureMain() {
	cooper := Person{"cooper" , 25 , true}
	log.Println(cooper.introduce())
}

type Person struct {
	Name string
	Age  int
	Sex  bool
}

func (p *Person) introduce() string {
	return fmt.Sprintf("I'm %v , %d years old , my sex is %v" , p.Name , p.Age , p.Sex)
}