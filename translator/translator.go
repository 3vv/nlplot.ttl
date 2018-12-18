// h 20181218
//
// Translator Interface

package translator

// Translator interface
type Translator interface {
	// Translate
	//   qry: `original text` dst: `target language` src: `source language`
	//   Result: `translated result` error: `result state`
	Translate(qry, dst, src string) (Result, error)
}

// NewTranslator
func NewTranslator(key string, sid string) *Base {
	return &Base{
		KEY: key,
		SID: sid}
}

// Base translator struct
type Base struct {
	API string // API URL
	QRY string // API query
	KEY string // API key
	SID string // Social ID
}

// Final translated result
type Result string
