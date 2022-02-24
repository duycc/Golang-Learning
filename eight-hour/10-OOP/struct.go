package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func main() {
	foo1()
}

func foo1() {
	var book1 Book
	book1.Title = "西游记"
	book1.Author = "吴承恩"

	book2 := Book{
		Title:  "三国演义",
		Author: "罗贯中",
	}

	fmt.Println(book1)
	fmt.Println(book2)

	setTitle(book1)
	setAuthon(&book2)

	fmt.Println(book1)
	fmt.Println(book2)
}

func setTitle(this Book) {
	this.Title = "红楼梦"
}

func setAuthon(this *Book) {
	this.Author = "施耐庵"
}
