package main

import (
	"fmt"
	"time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go sayHello() // start a new goroutine
    fmt.Println("Hello from main")

    // Wait for goroutine to finish (for demo only; use sync in real code)
    time.Sleep(time.Second)
}
