package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	for i := 0; i < 10; i++ {
		builder.WriteString("A")
	}

	fmt.Println(builder.String())
}
