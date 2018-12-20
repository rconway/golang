package main

import (
	"fmt"
	"log"
	"net/http"
)

var hostname = "0.0.0.0" // or could use e.g. ""
var port = 8080

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// NOT USED - but here is a most basic example of how to serve all files from the root dir down.
func serveRootDir() {
	http.Handle("/", http.FileServer(http.Dir(".")))
}

func servePublicDir() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

func main() {
	fmt.Println("FileServer started...")

	http.HandleFunc("/", serveIndex)
	http.Handle("/public/", servePublicDir())

	listenAddress := fmt.Sprintf("%s:%d", hostname, port)
	fmt.Printf("Listening on address: '%s'\n", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))

	fmt.Println("...FileServer completed.")
}
