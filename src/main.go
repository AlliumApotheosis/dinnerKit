package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Foxer! welcome to dinnerKit!")
	})
	http.ListenAndServe(":6336", nil)
}
