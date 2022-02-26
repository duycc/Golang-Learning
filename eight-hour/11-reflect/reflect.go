package main

import (
	"fmt"
	"reflect"
)

func main() {
	// foo1()
	foo2()
}

func foo1() {
	num := 1.2345
	reflectType(num)
}

func foo2() {
	user := User{1, "DuYong", 27}
	DoFileAndMethod(user)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("User is called...")
	fmt.Printf("%v\n", *this)
}

func DoFileAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType: ", inputType.Name())
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue: ", inputValue)

	// 通过Type获取里面的字段
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过Type获取方法
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}

func reflectType(param interface{}) {
	fmt.Println("Type: ", reflect.TypeOf(param))
	fmt.Println("Value: ", reflect.ValueOf(param))
}
