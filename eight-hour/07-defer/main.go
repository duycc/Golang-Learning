package main

import "fmt"

func main() {
	returnAndDefer()
	// returnFunc()...
	// deferFunc()...
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func returnFunc() int {
	fmt.Println("returnFunc()...")
	return 0
}

func deferFunc() int {
	fmt.Println("deferFunc()...")
	return 0
}
