package main

import "time"

func main() {
	// ch := make(chan int)
	//
	// go func() {
	// 	ch <- 5
	// }()
	//
	// value := <-ch
	// println(value)

	ch := make(chan int, 5)

	go func() {
		for i := range 10 {
			ch <- i
		}
	}()

	time.Sleep(time.Second * 5)
	for v := range ch {
		println(v)
	}
}
