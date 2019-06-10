package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/batman", batman)
	http.ListenAndServe(":5000", nil)
}

func batman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<image src="/resources/batman.jpg" />`)
}
