package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg2 sync.WaitGroup
)

func main(){
	wg2.Add(2)
	go incr()
	go incr()

	wg2.Wait()
	fmt.Printf("counter=%d\n", counter)

}

func incr(){
	defer wg2.Done()

	for i := 0; i < 2; i++ {
		//原子操作，同步值，同一时刻只有一个goroutine进行该操作
		atomic.AddInt64(&counter, 1)

		//多线程竞争，存在值覆盖情况
		//counter++
		//从当前线程退出，放回队列
		runtime.Gosched()
	}
}
