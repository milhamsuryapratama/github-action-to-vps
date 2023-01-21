package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there, I'm M. Ilham Surya Pratama")
	})

	http.ListenAndServe(":3000", nil)
}