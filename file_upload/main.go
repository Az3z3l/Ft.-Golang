package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	challid := r.FormValue("challenge_id")

	fmt.Println(challid)
	fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	folder := "challenges/" + challid + "/"
	fmt.Println(folder)
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("gere")
		os.MkdirAll(folder, 0700)
	}

	f, err := os.OpenFile(folder+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
	_, _ = io.Copy(f, file)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome home!")
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "./challenges", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/file", UploadFile).Methods("POST")
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(":2020", router))
}
