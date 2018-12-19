// h 20181219
//
// Command Line Interface for Youdao Translator

package main

import (
	"flag"
	"log"

	. ".."
)

// Command entry
func main() {
	flag.Parse()
	t := NewYoudaoTranslator(*key, *sid)
	r, e := t.Translate(*ori, *dst, *src)
	log.Printf("%v %v", e, r)
}

// Command flags
var (
	// User
	ori = flag.String("ori", "When you are old and grey and full of sleep, and nodding by the fire, take down this book, and slowly read.", "Original text") // Original text
	dst = flag.String("dst", "zh", "Target language")                                                                                                        // Target language
	src = flag.String("src", "auto", "Source language")                                                                                                      // Source language
	// Developer
	key = flag.String("key", "", "API key")   // API key
	sid = flag.String("sid", "", "Social ID") // Social ID
)
