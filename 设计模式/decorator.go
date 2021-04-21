package main

/**
意图: 装饰器模式动态地将责任附加到对象上。若要扩展功能，装饰者提供了比继承更有弹性的替代方案。
关键代码: 装饰器和被装饰对象实现同一个接口，装饰器中使用了被装饰对象
应用实例:
JAVA中的IO流
new DataInputStream(new FileInputStream("test.txt"));
人穿衣服，人则为被装饰对象，衣服为装饰器
 */
type Person interface {
	role()
}

type PersonDecorator struct {
	person Person
}

func (*PersonDecorator) role() {
}

//被装饰对象
type Worker struct {
}

func (*Worker) role() {
	println("i am worker")
}

type Poolman struct {
	PersonDecorator
}

func (p *Poolman) role() {
	p.person.role()
	println("pool man")
}

type Richman struct {
	PersonDecorator
}

func (p *Richman) role() {
	p.person.role()
	println("rich man")
}



