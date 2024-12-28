package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for timer...")
	<-timer.C // Block until timer sends a value on its channel
	fmt.Println("Timer expired!")
}
