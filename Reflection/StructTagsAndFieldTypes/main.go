package main

import (
	"fmt"
	"reflect"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}

    v := reflect.ValueOf(p)
    t := reflect.TypeOf(p)

    for i := 0; i < v.NumField(); i++ {
        fieldValue := v.Field(i)
        fieldType := t.Field(i)

        fmt.Printf("Field Name: %s\n", fieldType.Name)
        fmt.Printf("Field Type: %s\n", fieldValue.Type())
        fmt.Printf("Field Value: %v\n", fieldValue.Interface())

        // Struct tags
        tag := fieldType.Tag.Get("json")
        fmt.Printf("JSON Tag: %s\n\n", tag)
    }
}
