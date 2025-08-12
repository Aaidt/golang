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
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- "Error while reading body"
		return
	}

	p := post{}

	json.Unmarshal(body, &p)

	ch <- p.Title
}

func main() {
	ch := make(chan string)
	urls := [3]string{"https://jsonplaceholder.typicode.com/posts/1", "https://jsonplaceholder.typicode.com/posts/2", "https://jsonplaceholder.typicode.com/posts/3"}

	for _, url := range urls {
		go SendRequest(url, ch)
	}

	for range urls {
		select {
		case res := <-ch:
			{
				fmt.Println(res)
			}
		}
	}

	ans := <-ch

	fmt.Println(ans)
}
