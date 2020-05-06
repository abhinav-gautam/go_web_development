package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src='/toby.jpg'>
`)
}
func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
	}
	defer f.Close()

	io.Copy(w, f)
}
func main() {
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)

}
