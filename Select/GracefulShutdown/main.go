package main

import (
	"fmt"
)

func worker(ch chan int, quit chan bool) {
	for {
		select {
		case job := <-ch:
			fmt.Println("Processing job:", job)
		case <-quit:
			fmt.Println("Worker shutting down")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go worker(ch, quit)

	for i := 1; i <= 3; i++ {
		ch <- i
	}

	quit <- true
}
