package main

import (
	"fmt"
	"reflect"
)

func main() {
    numbers := []int{10, 20, 30}

    v := reflect.ValueOf(numbers)
    fmt.Println("Slice Kind:", v.Kind())

    // Modify an element
    elem := v.Index(1)
    if elem.CanSet() {
        elem.SetInt(99)
    }

    // Append a new element
    newValue := reflect.ValueOf(40)
    v = reflect.Append(v, newValue)

    // Convert back to interface
    updatedSlice := v.Interface().([]int)
    fmt.Println("Updated Slice:", updatedSlice)
}
