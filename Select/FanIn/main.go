package main

import (
	"fmt"
	"time"
)

func producer(ch chan string, name string, delay time.Duration) {
	for {
		time.Sleep(delay)
		ch <- fmt.Sprintf("Message from %s", name)
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go producer(ch1, "Producer1", 1*time.Second)
	go producer(ch2, "Producer2", 2*time.Second)

	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
