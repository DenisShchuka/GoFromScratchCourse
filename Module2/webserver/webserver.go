package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/gopher", gopherHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.ListenAndServe("localhost:8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You are ready to Go</h1>")
	fmt.Fprintf(w, "<a href='gopher'>What is Go gopher</a>")
}

func gopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go gopher</h1>")
	fmt.Fprintf(w, "<p>The Go gopher is an iconic macot of the Go project</p>")
	fmt.Fprintf(w, "<img src='img/gopher.jpg' />")
}
