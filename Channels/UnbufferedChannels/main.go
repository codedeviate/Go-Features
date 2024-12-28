package main

import (
	"fmt"
)

func sendData(ch chan int) {
	ch <- 10 // Send data
}

func main() {
	ch := make(chan int)

	go sendData(ch) // Start goroutine

	data := <-ch // Receive data
	fmt.Println("Received:", data)
}
