package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic" //原子操作库
)

//声明全局变量
var (
	counter   int64
	waitGroup = sync.WaitGroup{}
)
//使用原子锁锁住资源,防止并发操作同一个资源出现错误
func main() {

	waitGroup.Add(2)
	go incCounter()
	go incCounter()

	waitGroup.Wait()
	fmt.Println("final counter", counter)

}

func incCounter() {
	defer waitGroup.Done()
	for count := 0; count < 3; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}

}
