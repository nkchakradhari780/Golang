package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello again")
}
func sayHii() {
	fmt.Println("Hi")
}

func main() {
	fmt.Println("Learning goroutine")
	go sayHello()
	go sayHii()
	time.Sleep(3000 * time.Millisecond)

}