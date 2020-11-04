package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type test_struct struct {
	Email       string
	Uname       string
	Pword       string
	Age         string
	Institution string
	Contact     string
	Place       string
	District    string
	State       string
}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	fmt.Println(t.Uname)
	fmt.Println(t.Pword)
}

func main() {
	http.HandleFunc("/", parseGhPost)
	http.ListenAndServe(":2123", nil)
}
