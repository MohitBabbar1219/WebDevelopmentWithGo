package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":5000", nil)
}

func expire(w http.ResponseWriter, r *http.Request) {
	receivedCookie, err := r.Cookie("abyss")
	if err == http.ErrNoCookie {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}
	receivedCookie.MaxAge = -1 // delete cookie
	http.SetCookie(w, receivedCookie)
	http.Redirect(w, r, "/read", http.StatusSeeOther)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "abyss",
		Value: "madness",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "IN CHROME - GO TO: DEV TOOLS > APPLICATION > COOKIES")
}

func read(w http.ResponseWriter, r *http.Request) {
	receivedCookie, err := r.Cookie("abyss")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "Your cookie:", receivedCookie)
}
