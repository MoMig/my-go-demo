package main

/*
*命令源码文件：文件声明为main包，并且包含无参数声明无结果声明的函数
*对于独立程序，源码文件只有一个
*go run demo1.go --help 查看参数说明
 */

import (
	"flag"
	"fmt"
)

var name string
var age int

func init() {
	//参数值地址 参数名称 默认值 参数描述
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
	flag.IntVar(&age, "age", 18, "The greeting object's age.")
}

func main() {
	flag.Parse()
	fmt.Printf("hello %s\n", name)
	fmt.Println("age= ", age)
}
