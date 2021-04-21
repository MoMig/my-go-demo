package main

import (
	"flag"
	"fmt"
)

var key1 string
var key2 = "key2"

func main() {
	var key3 = "key3"
	println(key3)

	key4 := "key4"
	println(key4)

	// var name =右边的表达式，可以变为针对getTheFlag函数的调用表达式
	// Go 语言的类型推断可以明显提升程序的灵活性，使得代码重构变得更加容易，
	// 同时又不会给代码的维护带来额外负担（实际上，它恰恰可以避免散弹式的代码修改），更不会损失程序的运行效率
	var val = getVal()
	flag.Parse()
	fmt.Printf("%v", *val)

}

// 可以随意改变getTheFlag函数的内部实现，及其返回结果的类型，而不用修改main函数中的任何代码
func getVal() *string {
	return flag.String("name", "everyone", "The greeting object.")
}