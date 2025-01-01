package main

import "fmt"

type Writer interface {
	Write(data []byte) (int, error)
}

type Closer interface {
	Close() error
}

// Embedded interface
type WriteCloser interface {
	Writer
	Closer
}

type MyWriterCloser struct{}

func (m MyWriterCloser) Write(data []byte) (int, error) {
	return len(data), nil
}

func (m MyWriterCloser) Close() error {
	return nil
}

func main() {
	var wc WriteCloser = MyWriterCloser{}
	n, err := wc.Write([]byte("Hello"))
	fmt.Println("Wrote bytes:", n, "Error:", err)

	if err := wc.Close(); err != nil {
		fmt.Println("Close error:", err)
	}
}
