// h 20181224
//
// Command Line Interface for TTL Vendor

package main

import (
	"flag"
	"log"

	ttl "../.."
	_ "../baidu/vendor"
	_ "../youdao/vendor"
)

// Command entry
func main() {
	log.Println(ttl.Translate("baidu", *ori, *dst, *src))
	log.Println(ttl.Translate("youdao", *ori, *dst, *src))
}

func init() {
	flag.Parse()
}

// Command flags
var (
	ori = flag.String("ori", "When you are old and grey and full of sleep, and nodding by the fire, take down this book, and slowly read.", "Original text") // Original text
	dst = flag.String("dst", "zh", "Target language")                                                                                                        // Target language
	src = flag.String("src", "auto", "Source language")                                                                                                      // Source language
)
