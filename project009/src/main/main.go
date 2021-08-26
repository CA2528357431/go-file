// 数据序列化为json

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := struct {
		Name string
		Age int
		School []string
	}{
		Name:   "ca",
		Age:    18,
		School: []string{"hust","hit"},
	}
	// 记得把struct的参数变成public

	b := []int{1,2,3}

	c := make(map[string]bool)
	c["rich"] = true
	c["cool"] = true
	c["old"] = false

	d := "helloworld"


	ajson, err := json.Marshal(a)
	if err!=nil{
		println(err)
	}
	fmt.Println(string(ajson))

	bjson, err := json.Marshal(b)
	if err!=nil{
		println(err)
	}
	fmt.Println(string(bjson))

	cjson, err := json.Marshal(c)
	if err!=nil{
		println(err)
	}
	fmt.Println(string(cjson))

	djson, err := json.Marshal(d)
	if err!=nil{
		println(err)
	}
	fmt.Println(string(djson))



}


