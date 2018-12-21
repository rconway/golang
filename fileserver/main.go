package main

import (
	"fmt"
	"log"
	"net/http"
)

// Listening address parts.
var hostname = "0.0.0.0" // or could use e.g. ""
var port = 8080

// Serve the index.html file contents.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// FUNC NOT USED - but here is a most basic example of how to serve all files from the root dir down.
func serveRootDir() {
	http.Handle("/", http.FileServer(http.Dir(".")))
}

// Serve all files under the 'public' directory, stripping off the '/resources' prefix so they appear
// as if they are located from the root.
func servePublicDir() http.Handler {
	return http.StripPrefix("/resources/", http.FileServer(http.Dir("./public")))
}

func main() {
	fmt.Println("FileServer started...")

	// Routing
	http.HandleFunc("/", serveIndex)
	http.Handle("/resources/", servePublicDir())

	// Disable favicon
	//http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/favicon.ico")
	})

	// Start listening
	listenAddress := fmt.Sprintf("%s:%d", hostname, port)
	fmt.Printf("Listening on address: '%s'\n", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))

	// This will never get executed
	fmt.Println("...FileServer completed.")
}
