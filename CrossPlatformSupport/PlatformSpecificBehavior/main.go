package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Running on Windows: Special handling here.")
	} else {
		fmt.Println("Running on Unix-like system: Different handling here.")
	}
}
