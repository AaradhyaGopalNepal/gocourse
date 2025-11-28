package main

import (
	"fmt"
	dio "net/http"
)

func main() {
	fmt.Println("Hello go standard library")
	resp, err := dio.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status: ", resp.Status)
}
