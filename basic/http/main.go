package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println(k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello astaxie")

}

func sayG(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello astaxie")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	server := http.NewServeMux()
	server.HandleFunc("/h", sayG)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen", err)
	}
}
