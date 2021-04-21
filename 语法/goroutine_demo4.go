package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
* 互斥锁Mutex
* 创建临界区保证同一时间只有一个goroutine执行该片段代码
 */
var (
	counter4 int64
	wg4 sync.WaitGroup
	//定义代码临界区
	mutex sync.Mutex
)

func main(){
	wg4.Add(2)
	go incr4()
	go incr4()

	wg4.Wait()
	fmt.Printf("counter=%d\n", counter4)

}

func incr4(){
	defer wg4.Done()

	for i := 0; i < 2; i++ {
		mutex.Lock()
		//大括号为了让临界区看起来更清晰，不是必须的
		{
			value := counter4
			//当前线程退出并返回队列
			runtime.Gosched()
			value++
			counter4 = value
		}
		mutex.Unlock()
	}
}
