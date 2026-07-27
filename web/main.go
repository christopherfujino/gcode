package main

import (
	"fmt"
	"net/http"
)

func main() {
	var handler = http.FileServer(http.Dir("../docs"))
	const addr = "0.0.0.0:8080"
	fmt.Printf("Server listening at %s\n", addr)
	http.ListenAndServe(addr, handler)
}
