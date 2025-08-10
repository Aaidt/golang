package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	person := Person{Name: "Aadit", Age: 34, Email: ""}

	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
