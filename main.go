package main

import (
	"fmt"
	"time"
)

func expensiveOp(str string) {
	for i := range 30 {
		fmt.Println(str, "-", i)
	}
}

func main() {
	go expensiveOp("wassup")
	go expensiveOp("aadit-HP-Laptop-14s-cs3xxx")
	go expensiveOp("aadit")

	time.Sleep(time.Second * 2)
	fmt.Println("Done")
}
