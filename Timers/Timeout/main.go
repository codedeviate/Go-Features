package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.NewTimer(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Task completed before timeout.")
	case <-timeout.C:
		fmt.Println("Timeout occurred!")
	}
}
