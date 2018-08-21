package main

import (
	"fmt"
	"log"
	"net/http"
)

// HWhandler for somethign
func HWhandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Query().Get("uppercase") == "true" {
		fmt.Fprintf(w, "HELLO WORLD")
	} else if r.URL.Query().Get("uppercase") == "false" {
		fmt.Fprintf(w, "hello world")
	} else if r.URL.Query().Get("reverse") == "true" {
		fmt.Fprintf(w, "dlrow olleh")
	} else if r.URL.Query().Get("reverse") == "false" {
		fmt.Fprintf(w, "hello world")
	} else {
		fmt.Fprintf(w, "hello world")
	}
}

// Worldhandler for world string only
func Worldhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}

// Hellohandler for hello string only
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", HWhandler)
	http.HandleFunc("/hello", Hellohandler)
	http.HandleFunc("/world", Worldhandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
