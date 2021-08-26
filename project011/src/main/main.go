// json反序列化为数据

package main

import (

	"encoding/json"
	"fmt"
)

func main() {

	str1 := "{\"struct_name\":\"ca\",\"struct_age\":18,\"School\":[\"hust\",\"hit\"]}"
	var obj per

	err := json.Unmarshal([]byte(str1),&obj)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(obj)


	str2 := "[1,2,3]"
	li := make([]int,5)
	err = json.Unmarshal([]byte(str2),&li)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(li)

	str3 := "{\"cool\":true,\"old\":false,\"rich\":true}"
	dic := make(map[string]bool,5)
	err = json.Unmarshal([]byte(str3),&dic)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(dic)

	str4 := "\"helloworld\""
	var str string
	err = json.Unmarshal([]byte(str4),&str)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(str)


}

type per struct {
	Name string `json:"struct_name"`
	Age int `json:"struct_age"`
	School []string
}
