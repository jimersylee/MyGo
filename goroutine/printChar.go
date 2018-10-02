package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main()  {

	fmt.Println("Start Goroutines")
	//设置使用的cpu核心数
	runtime.GOMAXPROCS(1)
	fmt.Printf("CpuNum:%d\n",runtime.NumCPU())

	//设置waitGroup
	waitGroup:=sync.WaitGroup{}
	//设置2个goroutine
	waitGroup.Add(2)

	//创建一个匿名goroutine,打印小写字母表3次
	go func() {
		//在运行结束的时候通知
		defer waitGroup.Done()
		//循环打印3次
		for count:=0;count<3 ;count++  {
			//打印26个小写字母
			for char:='a';char<'a'+26;char++{
				fmt.Printf("%c",char)
				time.Sleep(1000)
			}
		}
	}()
	//创建一个匿名goroutine,打印26个大写字母3次
	go func() {
		//在运行结束的时候通知
		defer waitGroup.Done()
		//循环打印3次
		for count:=0;count<3 ;count++  {
			for char:='A';char<'A'+26;char++{
				fmt.Printf("%c",char)
				time.Sleep(1000)
			}
		}
	}()
	fmt.Println("Waiting to finish")

	waitGroup.Wait()
	fmt.Println("Work Done")

}