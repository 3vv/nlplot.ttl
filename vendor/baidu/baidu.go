// h 20181218
//
// Baidu Translator

package baidu

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	. "../../translator"
)

// Translate
func (t *BaiduTranslator) Translate(q, d, s string) (ret Result, err error) {
	for {
		// Prepare query string
		salt := time.Now().Format("20060102150405")
		sign := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s%s", t.SID, q, salt, t.KEY))))
		// http://api.fanyi.baidu.com/api/trans/product/apidoc#languageList
		if d == "" {
			d = "auto"
		}
		if s == "" {
			s = "auto"
		}
		// Request
		var rsp *http.Response
		rsp, err = http.Get(t.API + fmt.Sprintf(t.QRY, url.QueryEscape(q), d, s, t.SID, salt, sign))
		if err != nil {
			break
		}
		// Read response stream
		var msl []byte
		msl, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			break
		}
		// Unmarshal
		var mdl middleResult
		if err = json.Unmarshal(msl, &mdl); err != nil {
			break
		}
		// Unwrap
		if len(mdl.TargetResult) > 0 {
			ret, err = mdl.TargetResult[0].Dst, nil
		}
		// Finally
		if true {
			break
		}
	}
	return ret, err
}

// NewBaiduTranslator
func NewBaiduTranslator(k string, s string) *BaiduTranslator {
	t := NewTranslator(k, s)
	t.API = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	t.QRY = "?q=%s&to=%s&from=%s&appid=%s&salt=%s&sign=%s"
	return &BaiduTranslator{
		*t,
	}
}

// Translator
type BaiduTranslator struct {
	Base
}

// Result
type TargetResult struct {
	Dst Result `json:"dst"`
	Src string `json:"src"`
}
type middleResult struct {
	To           string         `json:"to"`
	From         string         `json:"from"`
	TargetResult []TargetResult `json:"trans_result"`
}
