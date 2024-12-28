package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	wg := sync.WaitGroup{}

	initialize := func() {
		fmt.Println("Initialization executed")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(initialize)
		}()
	}

	wg.Wait()
}
