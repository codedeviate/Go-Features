package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d running on %s\n", id, runtime.GOOS)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}
