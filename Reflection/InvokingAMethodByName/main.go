package main

import (
	"fmt"
	"reflect"
)

type Calculator struct{}

func (Calculator) Add(a, b int) int {
    return a + b
}

func main() {
    calc := Calculator{}
    calcValue := reflect.ValueOf(calc)

    // Retrieve method by name
    addMethod := calcValue.MethodByName("Add")
    if !addMethod.IsValid() {
        fmt.Println("Method not found")
        return
    }

    // Prepare arguments
    args := []reflect.Value{
        reflect.ValueOf(10),
        reflect.ValueOf(5),
    }

    // Call the method
    results := addMethod.Call(args)

    // Extract the result
    sum := results[0].Interface().(int)
    fmt.Println("Sum:", sum)
}
