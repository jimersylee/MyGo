package main

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Student struct {
	Human // 匿名字段
	school string
}


type Employee struct {
	Human // 匿名字段
	company string
}


//Human定义method
//首字母大写为public方法
func (h *Human) SayHi(){
	fmt.Printf("hi,this is Human's SayHi()")
}


//Employee的method重写Human的method

func (e *Employee) SayHi(){
	fmt.Println("hi,this is Employee's SayHi()")
}

func main(){
	mark:=Student{Human{"Mark",25,"222-2222"},"MIT"}
	sam:=Employee{Human{"sam",18,"1111-1111"},"Golang Inc"}

	mark.SayHi() // mark是student,继承了Human中的SayHi()method
	sam.SayHi()


}


