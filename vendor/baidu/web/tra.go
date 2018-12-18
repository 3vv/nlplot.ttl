// h 20181218
//
// RESTful implemention of Baidu Translator

package main

import (
	"io"
	"log"
	"net/http"
	"os"

	. ".."
)

// Trans
func Trans(w io.Writer, r *http.Request) (statusCode int) {
	qry := r.URL.Query()
	q := qry.Get("q")
	d := qry.Get("d")
	s := qry.Get("s")
	if f, e := t.Translate(q, d, s); e != nil {
		statusCode = http.StatusBadRequest
	} else {
		w.Write([]byte(f))
		statusCode = http.StatusOK
	}
	return statusCode
}

// init
func init() {
	// KEY & SID
	key, sid = os.Getenv("KEY"), os.Getenv("SID")
	// Assert $KEY and $SID
	if key == "" || sid == "" {
		log.Fatalln("Invalid $KEY/$SID")
	}
	// Init translator
	t = NewBaiduTranslator(key, sid)
}

// Variable
var (
	// Developer
	key string // API key
	sid string // Social ID
	// Translator
	t *BaiduTranslator
)
