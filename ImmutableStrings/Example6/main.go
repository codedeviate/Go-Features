package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is great"
	parts := strings.Split(str, " ")
	fmt.Println(parts) // ["Go", "is", "great"]
}
