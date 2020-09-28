package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHello)
	http.ListenAndServe(":8080", nil)
}

func handleHello(writer http.ResponseWriter, request *http.Request) {
	hello := "Hello, World!"
	fmt.Fprint(writer, hello)
}
