package interfacer
/*
 接口在同级别不定实例中使用
 接口类型接收指针类型
*/
import (
	"log"
	"math/rand"
	"time"
)

func interfacerMain() {
	cooper := Person{"cooper"}

	var personer Personer
	personer = &cooper
	personer.Say("fuck")
	personer.Do()
	personer.Walk()
	personer.Eat()

	//runrunrun()

	interfacerAssertions()

	interfacerTransform()

	interfacerQuery()

	polymorphism()
}

type Personer interface {
	Walk()
	Eat()
	Do()
	Say(string)
}

type Person struct {
	Name string
}

func (p *Person) Walk() {
	log.Printf("%v walk...", p.Name)
}

func (p *Person) Eat() {
	log.Printf("%v eat...", p.Name)
}

func (p *Person) Do() {
	log.Printf("%v do...", p.Name)
}

func (p *Person) Say(something string) {
	log.Printf("%v say : %v", p.Name , something)
}

// -----------------------------

type Animaler interface {
	Run()
}

type Human struct {}

func (h *Human) Run() {
	log.Println("Human Run...")
}

type Dog struct {}

func (d *Dog) Run() {
	log.Println("Dog Run...")
}

type Fish struct {}

func (f *Fish) Run() {
	log.Println("Fish Run...")
}

func runrunrun() {
	hChan := make(chan Human)
	dChan := make(chan Dog)
	fChan := make(chan Fish)

	rand.Seed(time.Now().UnixNano())
	go func() {
		t := rand.Intn(10)
		log.Println(t)
		<- time.After(time.Duration(t) * time.Second)
		h := Human{}
		hChan <- h
	}()

	go func() {
		t := rand.Intn(10)
		log.Println(t)
		<- time.After(time.Duration(t) * time.Second)
		d := Dog{}
		dChan <- d
	}()

	go func() {
		t := rand.Intn(10)
		log.Println(t)
		<- time.After(time.Duration(t) * time.Second)
		f := Fish{}
		fChan <- f
	}()

	var animaler Animaler
	select {
	case h := <-hChan:
		animaler = &h
	case d := <-dChan:
		animaler = &d
	case f := <-fChan:
		animaler = &f
	}

	animaler.Run()
}

func interfacerAssertions() {
	animals := map[string]interface{}{
		"dog":&Dog{},
		"fish":&Fish{},
	}
	for _ , obj := range animals {
		a , isA := obj.(Animaler)
		if isA {
			a.Run()
		}
	}
}

func interfacerTransform() {
	dog := new(Dog)

	var anim Animaler = dog

	dog2 := anim.(*Dog)

	dog2.Run()
}

func interfacerQuery() {
	fish := new(Fish)
	var anim Animaler = fish

	switch anim.(type) {
	case *Dog:
		log.Println("is a dog")
	case *Fish:
		log.Println("is a fish")
	}
}

type Mather interface {
	Count()
}

type PrimarySchool struct {}

func (ps *PrimarySchool) Count() {
	log.Println("primary student count 1 + 1")
}

type HighSchool struct {}

func (hs *HighSchool) Count() {
	log.Println("high student count 155 * 999")
}


func polymorphism() {

	f := func(school string) Mather {
		switch school {
		case "primary":
			return new(PrimarySchool)
		case "high":
			return new(HighSchool)
		}
		return nil
	}

	mather := f("primary")
	mather.Count()

}