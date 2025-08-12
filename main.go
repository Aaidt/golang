package main

import (
	"fmt"
)

type shape interface {
	area() float32
	perimeter() float32
}

type rect struct {
	length  float32
	breadth float32
}

type square struct {
	side float32
}

func (r rect) area() float32 {
	return r.length * r.breadth
}

func (r rect) perimeter() float32 {
	return 2 * (r.length + r.breadth)
}

func (s square) area() float32 {
	return s.side * s.side
}

func (s square) perimeter() float32 {
	return 4 * s.side
}

func calculateArea(s shape) {
	fmt.Println("The area is: ", s.area())
}

func calculatePerimeter(s shape) {
	fmt.Println("THe perimeter is: ", s.perimeter())
}

func main() {
	r := rect{length: 10, breadth: 20}
	sq := square{side: 30}

	calculateArea(sq)
	calculateArea(r)
	calculatePerimeter(sq)
	calculatePerimeter(r)
}
