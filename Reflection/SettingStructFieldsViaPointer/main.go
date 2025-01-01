package main

import (
	"fmt"
	"reflect"
)

type Config struct {
    Host string
    Port int
}

func main() {
    c := Config{Host: "localhost", Port: 8080}

    // Get addressable Value
    v := reflect.ValueOf(&c).Elem()

    // Modify the Host field
    hostField := v.FieldByName("Host")
    if hostField.CanSet() && hostField.Kind() == reflect.String {
        hostField.SetString("127.0.0.1")
    }

    // Modify the Port field
    portField := v.FieldByName("Port")
    if portField.CanSet() && portField.Kind() == reflect.Int {
        portField.SetInt(9090)
    }

    fmt.Printf("Updated Config: %+v\n", c)
}
