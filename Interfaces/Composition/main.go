package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Device struct{}

func (d Device) Read() string {
	return "Reading data"
}

func (d Device) Write(data string) {
	fmt.Println("Writing data:", data)
}

func main() {
	var rw ReadWriter = Device{}
	fmt.Println(rw.Read())
	rw.Write("GoLang")
}
