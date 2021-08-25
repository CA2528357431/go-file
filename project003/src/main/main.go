// 带缓冲区写入

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 打开
	filename := "./file/test.txt"
	file,err := os.OpenFile(filename,os.O_WRONLY|os.O_CREATE,0)
	if err != nil{
		fmt.Println(err)
	}
	// 参数分别是 文件路径、打开方式(见图)、win没锤子用的参数
	// 本例打开方式：只写、若文件不存在则创建

	// 写入：从头写入，后面有没有原数据不管
	// 追加：从末尾追加

	// 写入缓存
	writer := bufio.NewWriter(file)
	writer.Write([]byte("hello world\nhello golang"))
	// 写入文件
	err = writer.Flush()
	if err != nil{
		fmt.Println(err)
	}



	// 关闭
	err = file.Close()
	if err != nil{
		fmt.Println(err)
	}

}
