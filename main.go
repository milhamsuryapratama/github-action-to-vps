package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there, I'm M. Ilham Surya Pratama. I'm Software Engineer (Backend) at Ruangguru and Tukang Sapu at Gatherloop")
	})

	http.ListenAndServe(":3000", nil)
}
