// file_unix.go
//go:build linux || darwin
// +build linux darwin

package main

import "fmt"

func platformSpecificFunction() {
	fmt.Println("This code runs on Unix-like systems.")
}
