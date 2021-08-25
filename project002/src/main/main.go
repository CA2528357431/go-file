// 一次读取

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file := "./file/test.txt"
	li, err := ioutil.ReadFile(file)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(li))



}
