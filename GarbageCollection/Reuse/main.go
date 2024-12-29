package main

import "sync"

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) // Allocate 1KB slices
	},
}

func main() {
	buffer := pool.Get().([]byte)
	// Use buffer
	pool.Put(buffer) // Return buffer to the pool
}
