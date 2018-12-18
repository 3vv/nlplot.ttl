// h 20181218
//
// Command Line Interface for RESTful interface of Baidu Translator

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Runner
func main() {
	// Route
	r := mux.NewRouter()
	r.HandleFunc("/health", health)
	r.HandleFunc("/trans", trans)
	http.Handle("/", r)
	// Endpoint
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// KEY & SID
	log.Printf("Serving [host/port]%s:%s [key/sid]%t:%t, %s\n%s\n",
		host, port, key != "", sid != "", "press ^C to quit",
		"ENV:\n  HOST - host of endpoint\n  PORT - port of endpoint\n   KEY - API key\n   SID - Social ID")
	// Serve
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		panic(err)
	}
}
