package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", countVisits)
	http.ListenAndServe(":5000", nil)
}

func countVisits(w http.ResponseWriter, r *http.Request) {
	receivedCookie, err := r.Cookie("visitCount")
	if err == http.ErrNoCookie {
		receivedCookie = &http.Cookie{
			Name: "visitCount",
			Value: "0",
		}
	}
	count, err := strconv.Atoi(receivedCookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count += 1
	receivedCookie.Value = strconv.Itoa(count)

	http.SetCookie(w, receivedCookie)
	fmt.Fprintln(w, "You have visited this site", receivedCookie.Value, "times.")
}
