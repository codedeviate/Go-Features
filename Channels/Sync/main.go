package main

import (
	"fmt"
)

func worker(done chan bool) {
	fmt.Println("Working...")
	done <- true // Signal that work is done
}

func main() {
	done := make(chan bool)

	go worker(done)

	<-done // Wait for signal
	fmt.Println("Work completed")
}
