package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer expired!")
	}()

	time.Sleep(1 * time.Second)
	if timer.Stop() {
		fmt.Println("Timer stopped before expiration.")
	}
}
