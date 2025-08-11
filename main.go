package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {
	http.HandleFunc("/hello", Hello)

	http.ListenAndServe(":3000", nil)
}
