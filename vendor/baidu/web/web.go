// h 20181218
//
// RESTful interface of Baidu Translator

package main

import (
	"net/http"
)

// Health
func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// Trans
func trans(w http.ResponseWriter, r *http.Request) {
	Trans(w, r)
}
