package main

import (
	"fmt"
	"runtime"
	"sync"
)

//声明全局变量
var (
	counter   int64
	waitGroup sync.WaitGroup
	mutex     sync.Mutex
)
//使用互斥锁锁住资源,防止并发操作同一个资源出现错误
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
		mutex.Lock()
		value := counter
		runtime.Gosched()
		value++
		counter = value
		mutex.Unlock()
	}

}
