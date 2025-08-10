package main

import "sync"

func expensiveoperation() {
	sum := 0
	for i := range 100000 {
		sum += i
	}
	println(sum)
}

func main() {
	var w sync.WaitGroup

	for i := 0; i < 5; i++ {
		w.Add(1)

		go func() {
			defer w.Done()
			expensiveoperation()
		}()

	}
	w.Wait()
	println("Main function ended.")
}
