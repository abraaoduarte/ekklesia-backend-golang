package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("Hello world")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello router")
	})

	http.ListenAndServe(":80", nil)
}
