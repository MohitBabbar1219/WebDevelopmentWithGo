package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", batman)
	http.HandleFunc("/batman.jpg", batmanImage)
	http.ListenAndServe(":5000", nil)
}

func batmanImage(w http.ResponseWriter, r *http.Request) {
	image, err := os.Open("batman.jpg")
	if err != nil {
		log.Panic(err)
	}
	defer image.Close()

	imageInfo, err := image.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, r, image.Name(), imageInfo.ModTime(), image)
}

//func batmanImage(w http.ResponseWriter, r *http.Request) {
//	http.ServeFile(w, r, "batman.jpg")
//}

func batman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/batman.jpg" />`)
}
