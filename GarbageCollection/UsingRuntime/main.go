package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memStats runtime.MemStats

	// Force garbage collection
	runtime.GC()

	// Collect memory statistics
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Allocated Memory: %v KB\n", memStats.Alloc/1024)
	fmt.Printf("Number of GCs: %v\n", memStats.NumGC)
}
