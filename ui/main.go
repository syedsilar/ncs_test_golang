package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
	})
	http.ListenAndServe(":80", nil)
}
