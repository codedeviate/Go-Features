package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    rootCtx := context.Background()

    // Create a parent context with a 3-second timeout
    parentCtx, parentCancel := context.WithTimeout(rootCtx, 3*time.Second)
    defer parentCancel()

    // Create a child context with a 2-second timeout
    childCtx, childCancel := context.WithTimeout(parentCtx, 2*time.Second)
    defer childCancel()

    go func() {
        // Child will see Done if it hits 2s or if parent hits 3s first
        <-childCtx.Done()
        fmt.Println("Child context done:", childCtx.Err())
    }()

    go func() {
        <-parentCtx.Done()
        fmt.Println("Parent context done:", parentCtx.Err())
    }()

    time.Sleep(4 * time.Second)
    fmt.Println("Main: finished")
}
