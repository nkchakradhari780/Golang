package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web Request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)
	// fmt.Println("Response: ",res)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(data))
}
