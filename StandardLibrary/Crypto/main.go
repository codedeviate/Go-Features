package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	data := "password123"
	hash := md5.New()
	io.WriteString(hash, data)
	fmt.Printf("MD5 Hash: %x\n", hash.Sum(nil))
}
