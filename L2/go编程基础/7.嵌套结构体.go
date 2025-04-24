package main

import "fmt"

/*
	嵌套结构体
*/
type Person struct {
	name string
	age  int
}
type Student struct {
	person Person
	id     int
}

type A struct {
	Person
	Student
	a string
	b int
}

func main() {
	var p = Person{
		name: "sfx",
		age:  35,
	}
	fmt.Println(p)

	var s = Student{
		person: p,
		id:     1,
	}
	fmt.Println(s)
	var aa = A{
		Person:  p,
		Student: s,
		a:       "string",
		b:       10,
	}
	fmt.Println(aa)
}
