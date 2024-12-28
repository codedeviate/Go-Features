package main

import "fmt"

func sendMessage(ch chan string) {
	ch <- "Hello from goroutine"
}

func main() {
	ch := make(chan string)

	go sendMessage(ch)

	msg := <-ch // Receive message
	fmt.Println(msg)
}