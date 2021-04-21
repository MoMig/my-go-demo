package main

var block = "package"

func main(){
	println(block)

	block := "func"

	// 块代码
	{
		block := "inner"
		print(block)
	}

	println(block)

}
