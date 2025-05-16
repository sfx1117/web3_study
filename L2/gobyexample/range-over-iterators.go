package main

import (
	"fmt"
	"iter"
	"slices"
)

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	header *element[T]
	tail   *element[T]
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.header = &element[T]{val: v}
		lst.tail = lst.header
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) ALl() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.header; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}
func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.ALl() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.ALl())
	fmt.Println("all:", all)
	for n := range genFib() {

		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

//10
//13
//23
//all: [10 13 23]
//1
//1
//2
//3
//5
//8
