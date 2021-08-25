// 命令行参数

package main

import (
	"fmt"
	"os"
)

func main() {
	li := os.Args

	// li[0] 是文件名
	// 后面是参数
	// 参数都是string

	for _,v := range li{
		fmt.Printf("%v    %T\n",v,v)
	}

}
