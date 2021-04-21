package main

/**
* 接口的多态
* 1.方法(在func和方法名间添加一个类型参数)能给用户定义的类型添加新的行为
* 2.方法集定义了接口的接收规则，方法接收者为(t T),则参数值可为T或*T；方法接收者为(t *T),则参数值可为*T
 */

//定义接口
type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

//使用指针接收者实现notifier接口
func (u *user) notify(){
	println("user notify")
}

type admin struct {
	name string
	email string
}

//使用指针接收者实现notifier接口
func (a *admin) notify(){
	println("admin notify")
}

//接收notifier接口值，发送通知
func sendNotify( n notifier){
	n.notify()
}

func main() {
	u := user{
		name: "test1",
		email: "test1@qq.com",
	}
	sendNotify(&u)

	a := admin{
		name: "test2",
		email: "test2@qq.com",
	}
	sendNotify(&a)

}
