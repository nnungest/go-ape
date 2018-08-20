package main

import (
	"fmt"
	"log"
	"net/http"
)

func HWhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func Worldhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", HWhandler)
	http.HandleFunc("/hello", Hellohandler)
	http.HandleFunc("/world", Worldhandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
