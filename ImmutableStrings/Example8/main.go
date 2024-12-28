package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is great"
	fmt.Println(strings.Contains(str, "great")) // true
	fmt.Println(strings.Index(str, "is"))       // 3
}
