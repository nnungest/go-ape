package main

import (
	"fmt"
	"log"
	"net/http"
)

func HWhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
func Reversehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dlrow olleh")
}
func Upperhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")

}
func Worldhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world")
}
func Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func main() {
	http.HandleFunc("/", HWhandler)
	http.HandleFunc("/uppercase=false", HWhandler)
	http.HandleFunc("/uppercase=true", Upperhandler)
	http.HandleFunc("/reverse=false", HWhandler)
	http.HandleFunc("/reverse=true", Reversehandler)
	http.HandleFunc("/hello", Hellohandler)
	http.HandleFunc("/world", Worldhandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
