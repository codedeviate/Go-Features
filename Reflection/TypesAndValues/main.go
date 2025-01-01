package main

import (
	"fmt"
	"reflect"
)

func main() {
    var x int = 42
    v := reflect.ValueOf(x)
    t := reflect.TypeOf(x)

    fmt.Printf("Value: %v, Type: %v\n", v, t)
    fmt.Printf("Kind: %v\n", v.Kind())
}
