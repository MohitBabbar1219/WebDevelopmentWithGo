package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":5000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var fileContent string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		fileRead, fileHeader, err := r.FormFile("abyss")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer fileRead.Close()

		// info
		fmt.Println("\nfile:", fileRead, "\nfileHeader:", fileHeader, "\nerr:", err)

		// read
		bs, err := ioutil.ReadAll(fileRead)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileContent = string(bs)

		// write on server
		dst, err := os.Create(filepath.Join("./user/", fileHeader.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `
		<form method="post" enctype="multipart/form-data">
			<input type="file" name="abyss" />
			<input type="submit" />
		</form>
	` + fileContent)
}
