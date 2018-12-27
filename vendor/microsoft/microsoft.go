// h 20181226
//
// Microsoft Translator

package microsoft

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "../../translator"
)

// Translate
func (t *MicrosoftTranslator) Translate(q, d, s string) (ret Result, err error) {
	for {
		// Prepare query string
		// https://www.microsoft.com/en-us/translator/business/languages/
		// https://docs.microsoft.com/zh-cn/azure/cognitive-services/Translator/language-support
		// curl "https://api.cognitive.microsofttranslator.com/languages?api-version=3.0&scope=translation"
		if d == "" {
			d = "zh-Hans"
		}
		if s == "" {
			//s = ""
		}
		// Request
		var req *http.Request
		req, err = http.NewRequest("POST", t.API+fmt.Sprintf(t.QRY, d, s), strings.NewReader("[{\"Text\" : \""+q+"\"}]"))
		if err != nil {
			break
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
		req.Header.Add("Ocp-Apim-Subscription-Key", t.KEY)
		cli := &http.Client{Timeout: time.Second * 2}
		var rsp *http.Response
		rsp, err = cli.Do(req)
		if err != nil {
			break
		}
		// Read response stream
		defer rsp.Body.Close()
		var msl []byte
		msl, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			break
		}
		// Unmarshal
		var mdl interface{}
		err = json.Unmarshal(msl, &mdl)
		if err != nil {
			break
		}
		// Unwrap
		fmt.Printf("%#v\n", mdl)
		var tmp []byte
		tmp, err = json.MarshalIndent(mdl, "", "  ")
		if err != nil {
			break
		}
		ret = Bytes2Result(tmp)
		// Finally
		if true {
			break
		}
	}
	return ret, err
}

// NewMicrosoftTranslator
func NewMicrosoftTranslator(k string, s string) *MicrosoftTranslator {
	// KEY & SID
	// https://docs.azure.cn/zh-cn/articles/azure-operations-guide/cognitive-services/aog-cognitive-services-guidance
	// https://www.microsoft.com/cognitive-services/en-us/sign-up
	t := NewTranslator(k, s)
	// API & QRY
	t.API = "https://api.cognitive.microsofttranslator.com/translate"
	t.QRY = "?api-version=3.0&to=%s&from=%s"
	return &MicrosoftTranslator{
		*t,
	}
}

// Translator
type MicrosoftTranslator struct {
	Base
}

// Result
//type TargetResult Result
//type middleResult []byte
