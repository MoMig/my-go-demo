package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg3 sync.WaitGroup
)

func main(){
	wg3.Add(2)
	go doWork("tom")
	go doWork("jack")
	//休眠1秒
	time.Sleep(1 * time.Second)
	//线程安全 写值
	atomic.StoreInt64(&shutdown, 1)

	wg3.Wait()
	fmt.Printf("main end\n")

}

func doWork(p string){
	defer wg3.Done()

	for {
		fmt.Printf("%s is work\n", p)
		time.Sleep(200 * time.Millisecond)

		//线程安全的返回变量值
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("%s end\n", p)
			break
		}
	}
}
