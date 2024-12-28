package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Inside goroutine")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("Main function")
	time.Sleep(600 * time.Millisecond)
}
