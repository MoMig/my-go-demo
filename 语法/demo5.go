package main

import "fmt"

var container1 = []string{"zero", "one", "two"}

func main(){
	containter := map[int]string{0:"zero",1:"one",2:"two"}
	// 类型判断 等号右边为类型断言表达式
	value,ok := interface{}(container1).([]string)
	println(ok)
	println(value)
	fmt.Printf("%q \n",containter)
}