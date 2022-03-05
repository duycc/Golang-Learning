package main

import (
	"fmt"
	"time"
)

func foo1() {
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine end.")
		fmt.Println("goroutine start.")
		c <- 666
	}()

	num := <-c
	fmt.Println("num: ", num)
	fmt.Println("main goroutine end.")
}

func foo2() {
	// 带有缓冲的channel
	c := make(chan int, 3)
	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("sub goroutine end.")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("sub goroutine running, elem is: %d, len(c) = %d, cap(c) = %d\n", i, len(c), cap(c))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num is: ", num)
	}
	fmt.Println("main goroutine end.")
}

func foo3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			close(c)
		}
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main goroutine end.")
}

func foo4() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("main goroutine end.")
}

func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case data := <-quit:
			fmt.Println("quit: ", data)
			return
		}
	}
}

func main() {
	// foo1()
	// foo2()
	// foo3()
	// foo4()

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fib(c, quit)
}
