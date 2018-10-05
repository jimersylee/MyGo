package main

import (
	"fmt"
	"runtime"
	"sync"
)

//声明全局变量
var (
	counter   int
	waitGroup = sync.WaitGroup{}
)
//演示并发时的竞争导致的结果错误
func main() {

	waitGroup.Add(2)
	go incCounter()
	go incCounter()

	waitGroup.Wait()
	fmt.Println("final counter",counter)

}

func incCounter() {
	defer waitGroup.Done()
	value:=counter
	runtime.Gosched()
	value++
	counter=value
}
