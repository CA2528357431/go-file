package main

import (
	"flag"
	"fmt"
)

func main() {

	var name string
	var password string
	var id int
	var deleted bool
	var version float64

	// 分别是 目标变量、输入时的名称、默认值、说明
	flag.StringVar(&name,"n","","")
	flag.StringVar(&password,"p","","")
	flag.IntVar(&id,"i",0,"")
	flag.BoolVar(&deleted,"d",false,"")
	flag.Float64Var(&version,"v",0,"")
	// 布尔型只要输入就是真

	flag.Parse()

	fmt.Println(name,",")
	fmt.Println(password,",")
	fmt.Println(id,",")
	fmt.Println(deleted,",")
	fmt.Println(version,",")



}
