package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("."))) // will show all the files when we go to "/" route, but will show index.html if a file with that name is present
	http.HandleFunc("/batman", batman)
	http.ListenAndServe(":5000", nil)
}

func batman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<image src="/batman.jpg" />`)
}
