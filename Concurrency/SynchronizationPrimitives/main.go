package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (sc *SafeCounter) Increment() {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    sc.value++
}

func (sc *SafeCounter) Value() int {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    return sc.value
}

func main() {
    sc := &SafeCounter{}

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            sc.Increment()
        }()
    }
    wg.Wait()

    fmt.Println("Final counter value:", sc.Value())
}
