package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	header *element[T]
	tail   *element[T]
}

func (lst *List[T]) push(v T) {
	if lst.tail == nil {
		lst.header = &element[T]{val: v}
		lst.tail = lst.header
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.header; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")
	lst := List[int]{}
	lst.push(10)
	lst.push(13)
	lst.push(23)
	fmt.Println("list:", lst.AllElements())
}

//index of zoo: 2
//list: [10 13 23]
