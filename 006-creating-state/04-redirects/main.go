package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":5000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request method at foo:", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request method at bar:", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
	//http.Redirect(w, r, "/", http.StatusSeeOther)  // alternate way
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("your request method at barred:", r.Method)
	io.WriteString(w, `<h1>Barred</h1>`)
}
