package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Mohit-Key", "this is from mohit")
	//w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Out of all the things that you could have become, you chose to be a rat?</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":5000", d)
}
