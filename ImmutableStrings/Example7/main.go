package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := []string{"Go", "is", "great"}
	result := strings.Join(parts, " ")
	fmt.Println(result) // "Go is great"
}
