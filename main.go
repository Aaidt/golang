package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func SendRequest(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- "Error while sending request"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- "Error while reading body"
	}

	p := post{}

	json.Unmarshal(body, &p)

	ch <- p.Title
}

func main() {
	ch := make(chan string)
	go SendRequest("https://jsonplaceholder.typicode.com/posts/1", ch)

	ans := <-ch

	fmt.Println(ans)
}
