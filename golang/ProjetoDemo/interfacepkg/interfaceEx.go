package interfacepkg

import (
	"fmt"
	"math"
)

/* INTERFACES */

type geometry interface {
	area() float64
	perim() float64
}

type Rect struct {
	width, height float64
}

func (rect Rect) New(width, height float64) Rect {
	return Rect{width, height}
}

type Circle struct {
	radius float64
}

func (circ Circle) New(radius float64) Circle {
	return Circle{radius}
}

func (r Rect) area() float64 {
	return r.width * r.height
}
func (r Rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func Measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

/*
func main() {
    r := Rect{width: 3, height: 4}
    c := Circle{radius: 5}

    measure(r)
    measure(c)
} */
