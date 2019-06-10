package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Anything you want")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
