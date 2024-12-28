package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	queue := []int{}

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		mu.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		cond.Signal() // Notify waiting goroutine
		mu.Unlock()
	}

	for i := 0; i < 3; i++ {
		mu.Lock()
		for len(queue) == 2 {
			cond.Wait() // Wait until there's space in the queue
		}
		queue = append(queue, i)
		fmt.Println("Added to queue:", i)
		go removeFromQueue(1 * time.Second)
		mu.Unlock()
	}
	time.Sleep(2 * time.Second)
}
