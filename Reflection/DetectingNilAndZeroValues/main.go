package main

import (
	"fmt"
	"reflect"
)

func checkNil(value interface{}) {
    v := reflect.ValueOf(value)
    if !v.IsValid() {
        fmt.Println("Invalid (nil) value")
        return
    }

    // Check if it's nil
    if v.Kind() == reflect.Ptr && v.IsNil() {
        fmt.Println("Value is nil pointer")
        return
    }

    fmt.Println("Value is not nil")
}

func main() {
    var ptr *int
    checkNil(ptr)

    var num int = 42
    checkNil(num)

    // Also works with empty interface
    checkNil(nil)
}
