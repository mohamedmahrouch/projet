package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s ", r.URL.Path[1:])
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Fprintln(os.Stdout, "Server running on port 8088") // Message de confirmation
	err := http.ListenAndServe(":8088", nil )
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
}
