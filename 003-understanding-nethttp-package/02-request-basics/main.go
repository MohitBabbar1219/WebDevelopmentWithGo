package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Panic(err)
	}
	data := struct {
		Method      string
		Submissions url.Values
		URL         *url.URL
		Header      http.Header
	}{r.Method, r.Form, r.URL, r.Header}
	fmt.Println("Method:", data.Method)
	fmt.Println("URL:", data.URL)
	fmt.Println("Header:", data.Header)
	fmt.Println("Submissions:", data.Submissions)
	fmt.Println("Form value:", r.FormValue("fname"))
	_ = tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
