package main

import (
	"fmt"
)

type rect struct {
	width  int32
	height int32
}

func (r *rect) area() int32 {
	r.height = 30
	return r.width * r.height
}

func main() {
	r := rect{10, 20}

	fmt.Println(r.area())
	fmt.Println(r)
}
