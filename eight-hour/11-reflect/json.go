package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	foo()
}

func foo() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json Marshal error.", err)
		return
	}

	fmt.Println("jsonStr: ", jsonStr)

	myMovie := Movie{}
	if err = json.Unmarshal(jsonStr, &myMovie); err != nil {
		fmt.Println("json unmarshal error.", err)
		return
	}
	fmt.Println(myMovie)
}
