package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	address := "localhost:9000"
	fmt.Printf("Server started at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
	}
}