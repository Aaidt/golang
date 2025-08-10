package main

import "sync"

type safeCounter struct {
	m     sync.Mutex
	value int
}

func main() {
	var w sync.WaitGroup

	s := safeCounter{
		m:     sync.Mutex{},
		value: 0,
	}

	for i := range 100000 {
		w.Add(1)

		go func() {
			defer w.Done()
			defer s.m.Unlock()
			s.m.Lock()
			s.value += i
		}()
	}

	w.Wait()
	println(s.value)

	println("Main function ended.")
}
