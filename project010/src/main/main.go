// struct tag

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := struct {
		Name string `json:"struct_name"`
		Age int `json:"struct_age"`
		School []string
	}{
		Name:   "ca",
		Age:    18,
		School: []string{"hust","hit"},
	}
	// 若属性有json的tag，json的key会根据tag
	// tag 是反射的一种

	ajson, err := json.Marshal(a)
	if err!=nil{
		println(err)
	}
	fmt.Println(string(ajson))




}


