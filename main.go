package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	for range 2 {
		select {
		case first := <-ch1:
			fmt.Println(first)
		case second := <-ch2:
			fmt.Println(second)
		}
	}
}
