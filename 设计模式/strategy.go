package main

import "fmt"

/**
策略模式 (strategy)
意图: 定义一系列的算法,把它们一个个封装起来, 并且使它们可相互替换。
关键代码: 实现同一个接口
应用实例:
主题的更换，每个主题都是一种策略
旅行的出游方式，选择骑自行车、坐汽车，每一种旅行方式都是一个策略。
JAVA AWT 中的 LayoutManager。
 */

type IStrategy interface {
	printDocument(content string) (string,error)
}

type Printer struct {
	strategy IStrategy
}

func (o *Printer) setStrategy(strategy IStrategy) *Printer {
	o.strategy = strategy
	return o
}

func (o *Printer) todo(content string) {
	 r, err := o.strategy.printDocument(content)
	 if err != nil {
		 println(err)
		 return
	 }
	 println(r)
}

//具体策略
type PrintDoc struct {
}

func (*PrintDoc) printDocument(content string) (string,error) {
	println("print doc:" + content)
	r := fmt.Sprintf("len: %d", len(content))
	return r, nil
}

type PrintPdf struct {
}

func (*PrintPdf) printDocument(content string) (string,error) {
	println("print pdf:" + content)
	r := fmt.Sprintf("len: %d", len(content))
	return r, nil
}