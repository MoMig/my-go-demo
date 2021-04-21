package main

import "fmt"

type user2 struct {
	name string
	email string
}

func (u *user2) notify(){
	fmt.Printf("name=%s email=%s\n", u.name, u.email)
}

type admin2 struct {
	user2
	level string
}

func main(){
	u := &user2{
		name: "test1",
		email: "test1@qq.com",
	}
	u.notify()

	ad := admin2{
		user2: user2{
			name: "test2",
			email: "test2@qq.com",
		},
		level: "1",
	}
	//直接访问内部类型的方法
	ad.user2.notify()
	//内部类型的方法被提升到外部类型
	ad.notify()
}


