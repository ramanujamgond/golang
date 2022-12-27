// create a webserver in go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my home page\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}
