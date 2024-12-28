package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	stop := make(chan bool)

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t.Format("2006-01-02 15:04:05"))
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	stop <- true
	fmt.Println("Stopped ticking.")
}
