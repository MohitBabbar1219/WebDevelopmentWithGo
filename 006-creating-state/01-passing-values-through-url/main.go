package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("abyss")
	io.WriteString(w, "Did you just search for "+v)
}
