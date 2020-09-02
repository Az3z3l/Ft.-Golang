package main

import (
	"fmt"
	"net/http"
)

func helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

// func main() {
func dupe() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":5000", nil)
}
