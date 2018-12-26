// h 20181224
//
// Translator Vendor (Youdao)

package vendor

import (
	"log"
	"os"

	. ".."
	. "../../.."
)

// init
// To import only
func init() {
	// KEY & SID
	key, sid := os.Getenv("TTL_KEY_YOUDAO"), os.Getenv("TTL_SID_YOUDAO")
	// Check $KEY and $SID
	if key == "" || sid == "" {
		// Log $KEY and $SID
		log.Println("Nil $TTL_KEY_YOUDAO/$TTL_SID_YOUDAO")
	}
	// Register translator
	Register(
		// Name vendor
		"youdao",
		// Init translator
		NewYoudaoTranslator(key, sid))
}
