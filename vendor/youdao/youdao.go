// h 20181219
//
// Youdao Translator

package youdao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	. "../../translator"
)

// Translate
func (t *YoudaoTranslator) Translate(q, d, s string) (ret Result, err error) {
	for {
		// http://ai.youdao.com/docs/doc-trans-api.s#p05
		dat := url.Values{"i": {q}, "type": {fmt.Sprintf("%s2%s", s, d)}, "keyfrom": {"fanyi_web"}, "doctype": {"json"}, "xmlVersion": {"1.6"}, "ue": {"UTF-8"}, "typoResult": {"true"}, "flag": {"false"}}
		var rsp *http.Response
		rsp, err = http.PostForm(t.API+t.QRY, dat)
		if err != nil {
			break
		}
		defer rsp.Body.Close()
		var msl []byte
		msl, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			break
		}
		mdl := new(middleResult)
		json.Unmarshal(msl, mdl)
		ret, err = mdl.TranslateResult[0][0].Tgt, nil
		// Finally
		if true {
			break
		}
	}
	return ret, err
}

// NewYoudaoTranslator
func NewYoudaoTranslator(k string, s string) *YoudaoTranslator {
	t := NewTranslator(k, s)
	t.API = "http://fanyi.youdao.com/translate"
	t.QRY = "?" + url.Values{"smartresult": {"dict", "rule", "ugc"}}.Encode()
	return &YoudaoTranslator{*t}
}

// Translator
type YoudaoTranslator struct {
	Base
}

// Result
type TargetResult struct {
	Tgt Result `json:"tgt"`
	Src string `json:"src"`
}
type middleResult struct {
	TranslateResult [][]TargetResult `json:"translateResult"`
	ErrorCode       int              `json:"errorCode"`
	ElapsedTime     int              `json:"elapsedTime"`
	Type            string           `json:"type"`
}
