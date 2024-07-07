package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "POST" {
		http.Error(w, "method is not support", http.StatusNotFound)
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err %v", err)
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "form name:%s,address:%s", name, address)
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "method is not support", http.StatusNotFound)
	}
	fmt.Fprint(w, "hello")
}
func main() {
	fileser := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileser)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
