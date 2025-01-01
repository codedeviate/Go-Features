package main

import (
	"fmt"
)

func producer(ch chan<- int) {
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch)
}

func consumer(ch <-chan int, done chan<- bool) {
    for val := range ch {
        fmt.Println("Consumed:", val)
    }
    done <- true
}

func main() {
    ch := make(chan int)
    done := make(chan bool)

    go producer(ch)
    go consumer(ch, done)

    <-done // Wait for consumer to signal it's done
    fmt.Println("All values consumed.")
}
