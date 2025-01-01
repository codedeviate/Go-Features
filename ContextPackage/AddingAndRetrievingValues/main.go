package main

import (
	"context"
	"fmt"
)

type key string

const userIDKey key = "userID"

func main() {
    ctx := context.Background()

    // Store a user ID in the context
    ctx = context.WithValue(ctx, userIDKey, "12345")

    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    // Retrieve the user ID from context
    userID, ok := ctx.Value(userIDKey).(string)
    if !ok {
        fmt.Println("No user ID found in context")
        return
    }
    fmt.Println("processRequest: user ID is", userID)
}
