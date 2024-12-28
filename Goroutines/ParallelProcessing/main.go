package main

import (
	"fmt"
	"sync"
)

func compute(wg *sync.WaitGroup, id int) {
   defer wg.Done()
   fmt.Printf("Task %d started\n", id)
}

func main() {
   var wg sync.WaitGroup

   for i := 1; i <= 3; i++ {
	   wg.Add(1)
	   go compute(&wg, i)
   }

   wg.Wait()
   fmt.Println("All tasks completed")
}
