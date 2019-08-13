package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define my own middleware - see https://www.alexedwards.net/blog/making-and-using-middleware
func myMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[myMiddleware]: START")
		next.ServeHTTP(w, r)
		fmt.Println("[myMiddleware]: END")
	}
	return http.HandlerFunc(fn)
}

func main() {
	fmt.Println("Hello, world!")

	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /")
		http.SetCookie(w, &http.Cookie{Name: "c1", Value: "root2", Path: "/fred"})
		http.Redirect(w, r, "http://127.0.0.2:3000/larry", 307)
	})

	mux.Handle("/larry", http.RedirectHandler("http://127.0.0.3:3000/bob", 307))

	mux.HandleFunc("/bob", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /bob")
		http.SetCookie(w, &http.Cookie{Name: "c2", Value: "bob"})
		http.Redirect(w, r, "http://127.0.0.4:3000/fred", 307)
	})

	mux.HandleFunc("/fred", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit /fred")
		//w.WriteHeader(204)
		w.Write([]byte("<body><h6 style=\"text-align: center;\">TBD - result is here</h6></body>"))
	})

	// Using my middleware - see https://www.alexedwards.net/blog/making-and-using-middleware
	mux.Handle("/richard", myMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestURL := r.URL
		requestURI := r.RequestURI
		fmt.Println("URL=", r.URL)
		fmt.Println("requestURL=", requestURL.String())
		fmt.Println("requestURI=", requestURI)
		w.Write([]byte("this is richard"))
	})))

	// Serve some static content - the current directory under the url path '/public'.
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(""))))

	log.Fatal(http.ListenAndServe(":3000", mux))

	fmt.Println("EXITING...")
}
