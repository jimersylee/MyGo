package main

import "fmt"

//申明一个新的类型
type Person struct {
	name string
	age  int
}

func main() {
	var tom Person
	// 赋值初始化
	tom.name, tom.age = "Tom", 18
	//两个字段都写清除的初始化
	bob := Person{age: 15, name: "bob"}
	//按照struct定义的顺序初始化
	paul := Person{"paul", 19}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)

	//我们初始化一个学生
	mark:=Student{Human{"Mark",25,180},"Computer Science",222}
	//我们访问相应的字段
	fmt.Println("His name is ",mark.name) //访问的是Human中的name
	fmt.Println("His name is ",mark.Human.name) //访问的是Human中的name

}

func Older(p1, p2 Person) (Person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

//struct的匿名字段
type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human // 匿名自定义struct字段,那么默认的Student就包含了Human的所有字段
	speciality string
	int //匿名内置类型字段
}
