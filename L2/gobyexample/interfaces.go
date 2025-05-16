package main

import (
	"fmt"
	"math"
)

type geometry interface {
	aera() float64
	perim() float64
}
type rect struct {
	w float64
	h float64
}

func (r rect) aera() float64 {
	return r.w * r.h
}
func (r rect) perim() float64 {
	return 2 * (r.w + r.h)
}

type circle struct {
	r float64
}

func (c circle) aera() float64 {
	return math.Pi * c.r * c.r
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.r
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.aera())
	fmt.Println(g.perim())
}
func detectCircle(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.r)
	} else {
		fmt.Println("rect with w,h")
	}
}
func main() {
	r := rect{w: 3, h: 4}
	c := circle{r: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}

//{3 4}
//12
//14
//{5}
//78.53981633974483
//31.41592653589793
//rect with w,h
//circle with radius 5
