package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeChan(c, 666)
	time.Sleep(1 * time.Second)
}

func writeChan(c chan int, x int) {
	fmt.Println(x)
	c <- x // 阻塞
	close(c)
	fmt.Println(x)
}
