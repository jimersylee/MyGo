package main

import (
	"fmt"
	"runtime"
)

func say(s string){
	for i:=0;i<5;i++{
		runtime.Gosched() //runtime.GoSched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		fmt.Println(s)
	}
}


func main(){
	//goroutine test begin
	go say("world") //开一个新的goroutine执行
	say("hello") //当前goroutine执行
	//输出
	/*hello
	world
	hello
	world
	hello
	world
	hello
	world
	hello*/
	//goroutine test end


	//channels test begin

	a:=[]int{7,2,8,-9,4,0} //声明6个元素的
	c:=make(chan int) //创建channel c
	go sum(a[:len(a)/2],c) //计算前3个值的和
	go sum(a[len(a)/2:],c) //计算后3个值的和
	x,y:=<-c,<-c //从channel中获取值
	fmt.Println(x,y,x+y)



}



func sum(a []int,c chan int){
	total:=0
	for _,v:=range a{
		total+=v
	}
	c<-total // 将total的值赋予channel c
}




