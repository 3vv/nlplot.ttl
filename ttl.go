// h 20181221
//
// Text To Locale (language)

package ttl

import (
	"fmt"
	"sort"
	"sync"

	. "./translator"
)

// Register makes a translator available by the provided name
func Register(name string, translator Translator) {
	m.Lock()
	defer m.Unlock()
	for {
		if name == "" || translator == nil {
			break
		}
		if _, dup := vendors[name]; dup {
			panic("Register called twice or registered already for vendor: " + name)
			// ↗Prevent incorrect context if already registered by others
		}
		vendors[name] = translator
		// Finallly
		if true {
			break
		}
	}
}

// Unregister all the translators
func unregisterAllTranslators() {
	m.Lock()
	defer m.Unlock()
	vendors = make(map[string]Translator)
}

// Translate
//   vdr: `translator vendor`
//   qry: `original text` dst: `target language` src: `source language`
//   ret: `translated result` err: `result state`
func Translate(vdr, qry, dst, src string) (ret Result, err error) {
	m.RLock()
	vendor, ok := vendors[vdr]
	m.RUnlock()
	for {
		if !ok {
			err = fmt.Errorf("Unknown vendor: %q", vdr)
			break
		}
		ret, err = vendor.Translate(qry, dst, src)
		// Finallly
		if true {
			break
		}
	}
	return ret, err
}

// Translators returns a sorted list of the names of the registered vendors
func Translators() []string {
	m.RLock()
	defer m.RUnlock()
	var list []string
	for name := range vendors {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

// Variables
var (
	m       sync.RWMutex
	vendors = make(map[string]Translator)
)
