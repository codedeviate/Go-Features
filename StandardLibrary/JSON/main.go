package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Encode to JSON
	p := Person{Name: "Alice", Age: 30}
	jsonData, _ := json.Marshal(p)
	fmt.Println("JSON:", string(jsonData))

	// Decode from JSON
	var decoded Person
	json.Unmarshal(jsonData, &decoded)
	fmt.Printf("Decoded: %+v\n", decoded)
}
