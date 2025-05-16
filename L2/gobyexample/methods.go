package main

import "fmt"

type rect struct {
	w int
	h int
}

func (r *rect) area() int {
	return r.w * r.h
}
func (r rect) perim() int {
	return 2 * (r.w + r.h)
}
func main() {
	r := rect{w: 10, h: 5}
	fmt.Println(r.area())
	fmt.Println(r.perim())

	pr := &r
	fmt.Println(pr.area())
	fmt.Println(pr.perim())
}

//50
//30
//50
//30
