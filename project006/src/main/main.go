package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fileori,err := os.OpenFile("./file/ori.jpg",os.O_RDONLY,0)
	if err != nil{
		fmt.Println(err)
	}
	reader := bufio.NewReader(fileori)
	defer fileori.Close()

	filecopy,err := os.OpenFile("./file/copy.jpg",os.O_WRONLY|os.O_CREATE,0)
	if err != nil{
		fmt.Println(err)
	}
	writer := bufio.NewWriter(filecopy)
	defer filecopy.Close()

	_,err = io.Copy(writer,reader)

}
