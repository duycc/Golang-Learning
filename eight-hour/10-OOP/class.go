package main

import "fmt"

func main() {
	foo1()
}

type Human struct {
	Name string
	Sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human Eating...")
}

func (this *Human) Walk() {
	fmt.Println("Human Walking...")
}

type SupHuman struct {
	Human
	Level int
}

func (this *SupHuman) Fly() {
	fmt.Println("SupHuman Flieing...")
}

func (this *SupHuman) Walk() {
	fmt.Println("SupHuman Walking...")
}

func foo1() {
	h := Human{"LiBai", "male"}
	h.Eat()
	h.Walk()

	s := SupHuman{
		Human{"aLi", "female"},
		1,
	}

	s.Eat()
	s.Walk()
	s.Fly()
}
