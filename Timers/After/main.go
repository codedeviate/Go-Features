package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Waiting...")
	<-time.After(3 * time.Second) // Wait for 3 seconds
	fmt.Println("Time's up!")
}
