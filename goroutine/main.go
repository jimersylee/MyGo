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

	TestGoroutine()
	TestChannels()
	TestBufferedChannels()
	TestRangeAndClose()



}

/**

 */
func TestGoroutine(){
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
}

func sum(a []int,c chan int){
	total:=0
	for _,v:=range a{
		total+=v
	}
	c<-total // 将total的值赋予channel c
}

/**
阻塞型channel
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。所谓阻塞，也就是如果读取（value := <-ch）它将会被阻塞，直到有数据接收。其次，任何发送（ch<-5）将会被阻塞，直到数据被读出。无缓冲channel是在多个goroutine之间同步很棒的工具。


 */
func TestChannels(){
	//channels test begin
	a:=[]int{7,2,8,-9,4,0} //声明6个元素的
	c:=make(chan int) //创建channel c
	go sum(a[:len(a)/2],c) //计算前3个值的和
	go sum(a[len(a)/2:],c) //计算后3个值的和
	x,y:=<-c,<-c //从channel中获取值
	fmt.Println(x,y,x+y)
	//channels test end
}

/**
上面我们介绍了默认的非缓存类型的channel，不过Go也允许指定channel的缓冲大小，很简单，就是channel可以存储多少元素。ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。在这个channel 中，前4个元素可以无阻塞的写入。当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。
 */
func TestBufferedChannels(){
	//正确用法
	c := make(chan int, 2)//修改2为1就报错，修改2为3可以正常运行
	c <- 1
	c <- 2
	//错误写法
	/*d := make(chan int, 1)//修改2为1就报错，修改2为3可以正常运行
	//错误写法 c := make(chan int, 1)
	d <- 1
	d <- 2*/


	fmt.Println(<-c)
	fmt.Println(<-c)

}

/**
上面这个例子中，我们需要读取两次c，这样不是很方便，Go考虑到了这一点，所以也可以通过range，像操作slice或者map一样操作缓存类型的channel，请看下面的例子
 */
func TestRangeAndClose(){
	c:=make(chan int,10)
	go fibonacci(cap(c),c)
	for i:=range c{
		fmt.Println(i)
	}
}

func fibonacci(n int,c chan int){
	x,y:=1,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}
	close(c)
}




