package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)


func main() {
	fmt.Println("Web Request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n",res)
	// fmt.Println("Response: ",res)

	data,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(data))
}
 