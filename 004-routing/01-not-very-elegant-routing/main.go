package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "Dog")
	case "/cat":
		io.WriteString(w, "Cat")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":5000", d)
}
