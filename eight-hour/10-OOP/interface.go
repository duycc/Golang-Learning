package main

import "fmt"

// 多态现象，本质为一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat sleep...")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog sleep...")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(aniaml AnimalIF) {
	aniaml.Sleep()
	fmt.Println("Color: ", aniaml.GetColor())
	fmt.Println("Type: ", aniaml.GetType())
}

// interface{} 万能类型，可使用类型断言
func print(param interface{}) {
	fmt.Println("print()...")
	fmt.Println(param)

	if _, ok := param.(string); ok {
		fmt.Println("param is string type.")
	} else {
		fmt.Printf("param is %T type\n", param)
	}
}

func main() {
	// var aniaml AnimalIF
	// aniaml = &Cat{"Green"}
	// aniaml.Sleep()

	// aniaml = &Dog{"Yellow"}
	// aniaml.Sleep()

	// cat := Cat{"Green"}
	// dog := Dog{"Yellow"}

	// showAnimal(&cat)
	// showAnimal(&dog)

	print("hello world")
	print(32)
	print(22.99)
	print(Cat{"Green"})
}
