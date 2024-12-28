package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]string
}

func (sm *SafeMap) Write(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Read(key string) string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.data[key]
}

func main() {
	sm := SafeMap{data: make(map[string]string)}
	wg := sync.WaitGroup{}

	// Writing to the map
	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.Write("name", "GoLang")
	}()

	// Reading from the map
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Name:", sm.Read("name"))
	}()

	wg.Wait()
}
