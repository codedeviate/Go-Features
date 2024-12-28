package main

import (
	"fmt"
	"time"
)

func retryTask(maxRetries int, delay time.Duration) {
	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Attempt %d...\n", i)
		success := attemptTask()
		if success {
			fmt.Println("Task succeeded!")
			return
		}
		if i < maxRetries {
			time.Sleep(delay) // Wait before retrying
		}
	}
	fmt.Println("Task failed after retries.")
}

func attemptTask() bool {
	// Simulate a task with 50% success rate
	return time.Now().Unix()%2 == 0
}

func main() {
	retryTask(5, 1500*time.Millisecond)
}
