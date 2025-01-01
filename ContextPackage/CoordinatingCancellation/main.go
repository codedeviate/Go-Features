package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    for i := 0; i < 3; i++ {
        go worker(ctx, i)
    }

    // Wait for a while
    time.Sleep(5 * time.Second)
    fmt.Println("Main: all done")
}

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d: done, reason: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d: working...\n", id)
            time.Sleep(1 * time.Second)
        }
    }
}
