package main // 包名，main包的main函数是程序入口

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello, go...")
	time.Sleep(time.Second)
}

// 运行方式：
// 1. go build hello.go && ./hello
// 2. go run hello.go
