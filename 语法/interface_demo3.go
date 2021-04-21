package main

import "fmt"

type notifier3 interface {
	notify()
}

type user3 struct {
	name string
	email string
}

func (u *user3) notify(){
	fmt.Printf("[user] name=%s email=%s\n", u.name, u.email)
}

type admin3 struct {
	user3
	level string
}

func (u *admin3) notify(){
	fmt.Printf("[admin] name=%s email=%s\n", u.name, u.email)
}

func sendNotity3(n notifier3){
	n.notify()
}

func main(){
	ad := admin3{
		user3: user3{
			name: "test3",
			email: "test3@qq.com",
		},
		level: "1",
	}
	//直接访问内部类型的方法
	ad.user3.notify()
	//内部类型的方法没有被提升到外部类型
	ad.notify()
}


