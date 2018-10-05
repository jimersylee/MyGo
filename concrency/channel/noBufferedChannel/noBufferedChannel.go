package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

/**
初始化函数,设置随机数种子
 */
func init() {
	rand.Seed(time.Now().UnixNano())
}

/**
使用无缓冲通道
这是示例使用无缓冲通道模拟两个goroutine之间的网球比赛
 */
func main() {
	court := make(chan int)
	wg.Add(2)
	//启动两个选手
	go player("A", court)
	go player("B", court)
	//发球
	court <- 1
	//等待结束
	wg.Wait()
	fmt.Println("finished")

}

/**
模拟运动员击球
 */
func player(name string, court chan int) {
	//通知执行结束
	defer wg.Done()
	//从通道中获取球
	for {
		ball, ok := <-court
		if !ok {
			//说明对手丢球了,把通道关了,则我赢了
			fmt.Println(name, "win!")
			return
		}

		//还没结束,继续玩,使用随机数来决定是否丢球
		num := rand.Intn(100)
		if num%13 == 0 {
			//丢球啦,关闭通道
			fmt.Println(name," lose the ball")
			close(court)
			return
		}


		//成功回球
		fmt.Println(name, " hit ", ball, " times")
		ball++
		court <- ball

	}

}
