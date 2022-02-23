package main

import "fmt"

func swapByValue(x int, y int) {
	tmp := x
	x = y
	y = tmp
}

func swapByPtr(px *int, py *int) {
	tmp := *px
	*px = *py
	*py = tmp
}

func main() {
	x, y := 10, 20
	swapByValue(x, y) // x: 10, y: 20
	swapByPtr(&x, &y) // x: 20, y: 10

	fmt.Printf("x: %d, y: %d\n", x, y)
}
