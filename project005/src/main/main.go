package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Stat("./file/test.txt")
	if err==nil{
		fmt.Println("exist")
	}else if err==os.ErrNotExist {
		fmt.Println("not exist")
	}else {
		fmt.Println("error")
	}

}
