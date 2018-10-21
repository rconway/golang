package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, world!")

	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "c1", Value: "root2", Path: "/fred"})
		http.Redirect(w, r, "/larry", 301)
	})

	mux.Handle("/larry", http.RedirectHandler("/bob", 302))

	mux.HandleFunc("/bob", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "c2", Value: "bob"})
		http.Redirect(w, r, "/fred", 303)
	})

	mux.HandleFunc("/fred", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
		w.Write([]byte("this is fred"))
	})

	http.ListenAndServe(":3000", mux)
}
