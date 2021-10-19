package main

import "fmt"

func main() {
	rect := rectangle{
		length: 20, breadth: 30,
	}
	area := rect.area()
	fmt.Println(area)
}

type rectangle struct {
	length, breadth int
}

func (r rectangle) area() int {
	return r.length * r.breadth
}
