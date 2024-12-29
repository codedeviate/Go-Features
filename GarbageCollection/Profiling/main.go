package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.ListenAndServe("localhost:6060", nil)
}
