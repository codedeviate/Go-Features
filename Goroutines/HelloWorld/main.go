package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sayHello() // Start the goroutine
	for i := 0; i < 5; i++ {
		fmt.Println("World")
		time.Sleep(200 * time.Millisecond)
	}
}
