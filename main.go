package main

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
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}()

	val, ok := <-ch
	println(val, ok)

	for v := range ch {
		println(v)
	}
}
