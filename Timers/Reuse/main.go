package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(1 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("First timer expired!")
	}()

	time.Sleep(2 * time.Second)

	timer.Reset(3 * time.Second) // Reset the timer
	// <-timer.C
	fmt.Println("Second timer expired!")
}
