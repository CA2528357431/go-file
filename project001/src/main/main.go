// 带缓冲区分次数读取

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开
	filename := "./file/test.txt"
	file,err := os.OpenFile(filename,os.O_RDONLY|os.O_CREATE,0)
	if err != nil{
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for{
		// 读到\n(包括\n)结束
		str,err := reader.ReadString('\n')

		// 如果文件结束则break
		if err==io.EOF {
			break
		}
		fmt.Println(str)

	}


	// 关闭
	err = file.Close()
	if err != nil{
		fmt.Println(err)
	}

}
