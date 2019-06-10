package main

import (
	"io"
	"net/http"
)

type hotdog int

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dog")
}

type hotcat int

func (hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cat")
}

func main() {
	//--------- Explicit servemux ---------
	//var d hotdog
	//var c hotcat
	//mux := http.NewServeMux()
	//mux.Handle("/dog/", d)
	//mux.Handle("/cat", c)
	//http.ListenAndServe(":5000", mux)
	//-------------------------------------

	//--------- Default servemux (better) ---------
	var d hotdog
	var c hotcat
	http.Handle("/dog/", d) // default servemux
	http.Handle("/cat", c)
	http.ListenAndServe(":5000", nil) // notice nil
	//---------------------------------------------

}
