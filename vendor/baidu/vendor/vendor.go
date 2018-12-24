// h 20181224
//
// Translator Vendor (Baidu)

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
	key, sid := os.Getenv("TTL_KEY_BAIDU"), os.Getenv("TTL_SID_BAIDU")
	// Check $KEY and $SID
	if key == "" || sid == "" {
		// Assert $KEY and $SID
		log.Fatalln("Invalid $KEY/$SID")
	}
	// Register translator
	Register(
		// Name vendor
		"baidu",
		// Init translator
		NewBaiduTranslator(key, sid))
}
