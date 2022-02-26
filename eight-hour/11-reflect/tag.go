package main

import (
	"fmt"
	"reflect"
)

func main() {
	re := Resume{}
	parseTag(re)
}

type Resume struct {
	Name string `info:"name" doc:"名字"`
	Sex  string `info:"sex"`
}

func parseTag(param interface{}) {
	t := reflect.TypeOf(param)

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", tagInfo, ", doc: ", tagDoc)
	}
}
