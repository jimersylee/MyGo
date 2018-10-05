package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
//10个任务
const TASK_NUM=10
//4个工人
const WORKER_NUB=4
/**
使用缓冲通道模拟4个工人处理10个任务
 */
func main(){
	//创建一个缓冲通道
	tasks:=make(chan int,TASK_NUM)
	//设置等待的goroutine数
	wg.Add(WORKER_NUB)
	//启动N个工作者
	for count:=0;count<WORKER_NUB ;count++  {
		go worker(tasks,count)
	}

	//分配工作
	for count:=0;count<TASK_NUM;count++{
		tasks<-count
	}

	//工作分配完了,可以关闭通道了(通道特性,关闭通道后,无法写入,但是还是取出剩余的)
	close(tasks)
	wg.Wait()
}


func worker(tasks chan int,worker int){
	defer wg.Done()
	for {
		//取出工作
		task, ok := <-tasks
		if !ok {
			//所有工作都做完啦
			fmt.Println("Worker",worker,":no work,shutting down")
			return
		}
		//还有工作
		fmt.Println("Worker", worker, " starts do task ", task)
		//处理工作,模拟处理过程
		sleepTime := rand.Intn(100)
		time.Sleep(time.Duration(sleepTime) * time.Microsecond)
		fmt.Println("Worker", worker, " has finished task ", task)
	}
}



