package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":4444", nil)
}
