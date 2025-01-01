package main

import (
	"fmt"
)

func main() {
    err := doOperation()
    if err != nil {
        fmt.Println("Operation failed with error:", err)
    } else {
        fmt.Println("Operation succeeded")
    }
}

func doOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("doOperation panic: %v", r)
        }
    }()

    riskyOperation()
    return nil
}

func riskyOperation() {
    panic("some unexpected condition")
}
