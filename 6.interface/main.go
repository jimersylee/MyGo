package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	money float32
}

/**
定义interface
 */

type Man interface {
	SayHi()//说hello
	Sing(lyrics string)//唱歌
}

type YoungChap interface {
	SayHi()
	Sing()
	BorrowMoney(amount float32)//借钱
}

type ElderlyGent interface{
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)//花费薪水
}





func (h Human) SayHi(){
	fmt.Printf("my nameis %s,say hi\n",h.name)
}
func (h Human) Sing(song string){
	fmt.Printf("my name is %s, sing\n",h.name)
}




func main(){
	mike:=Student{Human{"mike",11,"2231231"},"MIT",1111}
	sam:=Employee{Human{"sam",22,"21232"},"Google",11111}

	//定义Man类型的变量
	var man Man
	man=mike
	man.SayHi()
	man.Sing("lalala")

	man=sam
	man.SayHi()
	man.Sing("ahahha")




}
