package main

import (
	"fmt"
	"sync"
)

func printString(wg *sync.WaitGroup, str string) {
	defer wg.Done()
	fmt.Println(str)
}

func main() {
	var wg sync.WaitGroup
	str := "Immutable"

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printString(&wg, str)
	}

	wg.Wait()
}
