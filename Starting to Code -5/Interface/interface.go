package main

import (
	"fmt"
	"math"
)

func main() {
	var r geomentey = rectangle{width: 20, height: 20}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
	var c geomentey = circle{radius: 6}
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
}

type geomentey interface { //interface
	area() float64
	perimeter() float64
}
type rectangle struct { //struct
	height, width float64
}

func (r rectangle) area() float64 { //methods
	return r.height * r.width
}
func (r rectangle) perimeter() float64 {
	return 2*r.height + 2*r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
