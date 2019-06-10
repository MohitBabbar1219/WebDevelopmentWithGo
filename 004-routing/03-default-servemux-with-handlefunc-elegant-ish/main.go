package main

import (
	"io"
	"net/http"
)

func serveDog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

func serveCat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat")
}

func main() {
	http.HandleFunc("/dog/", serveDog)
	http.HandleFunc("/cat", serveCat)
	// This is not implementing the handler interface, do not get confused.
	// This is just a function with similar signature as handler interface
	http.ListenAndServe(":5000", nil)
}
