package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "Message from ch1"
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			ch2 <- "Message from ch2"
			time.Sleep(2 * time.Second)
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
