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
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<form metho="post">
		<input type="text" name="abyss" />
		<input type="submit">
</form><br />` + v)
}
