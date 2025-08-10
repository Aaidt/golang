package main

func expensiveoperation(ch chan<- bool) {
	sum := 0
	for i := range 100000000000 {
		sum += i
	}
	println(sum)
	ch <- true
}

func main() {
	ch := make(chan bool)

	go expensiveoperation(ch)
	<-ch
	println("Main function ended.")
}
