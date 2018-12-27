// h 20181227
//
// Translator Vendor (Microsoft)

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
	key, sid := os.Getenv("TTL_KEY_MICROSOFT"), os.Getenv("TTL_SID_MICROSOFT")
	// Check $KEY and $SID
	if key == "" || sid == "" {
		// Assert $KEY and $SID
		log.Fatalln("Invalid $TTL_KEY_MICROSOFT/$TTL_SID_MICROSOFT")
	}
	// Register translator
	Register(
		// Name vendor
		"microsoft",
		// Init translator
		NewMicrosoftTranslator(key, sid))
}
