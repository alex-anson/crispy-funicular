package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handles requests to the root URL.
func homePage(w http.ResponseWriter, r *http.Request) {
	// func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// Fprintf formats according to a format specifier and writes to w. It returns
	// the number of bytes written and any write error encountered.
	fmt.Fprintf(w, "Welcome to the homepage of my first Go repo 🂡")
	fmt.Println("Endpoint hit: homepage")
}

// Matches the URL path hit with a defined function
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

// Obvi most important ƒn ✨
func main() {
	handleRequests()
}
