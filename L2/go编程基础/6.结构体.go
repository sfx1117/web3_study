package main

import "fmt"

/*
	结构体
*/
//方式1
type Person struct {
	name string
	age  int
}

//方式2
type Person1 struct {
	name, age int
}

//方式3
type Person3 struct {
	name string `json:"name" gorm:"column: NAME"`
	age  int    `json:"age" gorm:"column:AGE"`
}

//方式4 匿名字段
type Person4 struct {
	string
	int
}

//匿名结构体 全局变量
var Person5 = struct {
	name string
	age  int
}{
	name: "sfx",
	age:  35,
}

func main() {
	//匿名结构体 局部变量
	var Person6 = struct {
		name string
		age  int
	}{
		name: "sfx",
		age:  35,
	}
	fmt.Println(Person6)
	fmt.Println(Person6.name)
	fmt.Println(Person6.age)
}
