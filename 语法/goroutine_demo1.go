package main
/**
* goroutine并发执行
 */

import (
	"fmt"
	"runtime"
	"sync"
)

//同步计数器 等待程序执行完成
var wg sync.WaitGroup

func main() {
	// 分配一个逻辑处理器
	runtime.GOMAXPROCS(1)
	//计数器设置为2，需等待2个goroutine执行完
	wg.Add(2)
	go printStr("A")
	go printStr("B")
	wg.Wait()
	fmt.Println("end")


}

func printStr(pre string) {
	defer wg.Done()
	for i := 2; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i % j == 0 {
				fmt.Printf("%s - %d\n", pre, i)
			}
		}
	}
}
