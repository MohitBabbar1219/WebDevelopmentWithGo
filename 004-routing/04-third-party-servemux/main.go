package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Deal(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Deal")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/deal", Deal)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":5000", router))
}
