// 一次写入

package main

import (
	"fmt"
	"os"
)

func main() {

	// 打开
	filename := "./file/test.txt"
	err := os.WriteFile(filename,[]byte("hello world"),0)
	if err != nil{
		fmt.Println(err)
	}


}