package main

import "fmt"

/*
	1. 不同长度数组是不同的类型
	2. 数组传参是值传递，会发生拷贝
	3. slice传参是引用传递
*/

func main() {
	// foo1()
	// foo2()
	// foo3()
	foo4()
}

func foo4() {
	s := []int{1, 2, 3} // len = 3, cap = 3, [1,2,3]
	// [0, 2)
	s1 := s[0:2] // [1, 2]
	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	// copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3) // s2 = [0,0,0]

	// 将s中的值依次拷贝到s2中
	copy(s2, s)
	fmt.Println(s2)
}

func foo3() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素1, numbers len = 4， [0,0,0,1], cap = 5
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向numbers切片追加一个元素2, numbers len = 5， [0,0,0,1,2], cap = 5
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// 向一个容量cap已经满的slice 追加元素，
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}

func foo2() {
	slice := []int{1, 2, 3, 4} // 动态数组，切片 slice
	fmt.Printf("slice type is %T\n", slice)
	printSlice(slice)
	for _, value := range slice {
		fmt.Println("value = ", value)
	}
}

func foo1() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray(myArray3)
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

func printArray(array [4]int) {
	// 值拷贝
	for index, value := range array {
		fmt.Println("index = ", index, ", value = ", value)
	}
	array[0] = 111
}

func printSlice(slice []int) {
	// 引用传递
	// _ 表示匿名的变量
	for _, value := range slice {
		fmt.Println("value = ", value)
	}
	slice[0] = 100
}
