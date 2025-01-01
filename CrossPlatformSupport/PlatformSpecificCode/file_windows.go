// file_windows.go
//go:build windows
// +build windows

package main

import "fmt"

func platformSpecificFunction() {
	fmt.Println("This code runs on Windows.")
}
