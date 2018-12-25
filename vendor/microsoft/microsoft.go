// h 20181225
//
// Microsoft Translator

package microsoft

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	. "../../translator"
)

// Translate
func (t *MicrosoftTranslator) Translate(q, d, s string) (ret Result, err error) {
	for {
		// Check token
		if t.lessToken(t.auth.ttl) {
			err = t.takeToken()
			if err != nil {
				break
			}
		}
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
		req, _ := http.NewRequest("GET", t.API+fmt.Sprintf(t.QRY, url.QueryEscape(q), d, s), nil)
		req.Header.Add("Authorization", "Bearer "+t.auth.tkn)
		cli := &http.Client{}
		var rsp *http.Response
		rsp, err = cli.Do(req)
		if err != nil {
			break
		}
		// Read response stream
		defer rsp.Body.Close()
		var mdl middleResult
		mdl, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			break
		}
		// Pick
		ret, err = Bytes2Result(String2Bytes(pickResult(Bytes2String(mdl)))), nil
		// Finally
		if true {
			break
		}
	}
	return ret, err
}

const lessHold = 2 * 1e9

func (t *MicrosoftTranslator) lessToken(timeout time.Duration) (yes bool) {
	if timeout > 0 {
		yes = t.tat.Add(timeout + lessHold).After(nowFunc())
	}
	return yes
}

var nowFunc = time.Now

func (t *MicrosoftTranslator) takeToken() (err error) {
	for {
		req, _ := http.NewRequest("POST", t.auth.URL, strings.NewReader(t.auth.VAL))
		cli := &http.Client{}
		var rsp *http.Response
		rsp, err = cli.Do(req)
		if err != nil {
			break
		}
		defer rsp.Body.Close()
		var msl []byte
		msl, err = ioutil.ReadAll(rsp.Body)
		if err != nil {
			break
		}
		mdl := make(map[string]interface{}, 4)
		err = json.Unmarshal(msl, &mdl)
		if err != nil {
			break
		}
		if token, ok := mdl["access_token"]; !ok {
			err = errors.New("")
			break
		} else {
			t.tkn, err = token.(string), nil
			t.tat = nowFunc()
			if expires, ok := mdl["expires_in"]; ok {
				t.ttl = expires.(time.Duration) * 1e9
			}
		}
		if true {
			break
		}
	}
	return err
}

// NewMicrosoftTranslator
func NewMicrosoftTranslator(k string, s string) *MicrosoftTranslator {
	t := NewTranslator(k, s)
	t.API = "http://api.microsofttranslator.com/v2/Http.svc/Translate"
	t.QRY = "?text=%s&to=%s&from=%s"
	a := &auth{}
	a.URL = "https://datamarket.accesscontrol.windows.net/v2/OAuth2-13"
	a.VAL = makeAuth(k, s)
	a.tat = nowFunc()
	a.ttl = 1
	return &MicrosoftTranslator{
		Base: *t,
		auth: *a,
	}
}

var regexpResult = regexp.MustCompile(`^<string xmlns=".+?">(.*)</string>$`)

func pickResult(result string) (ret string) {
	mat := regexpResult.FindStringSubmatch(result)
	if len(mat) == 2 {
		ret = mat[1]
	}
	return ret
}

func makeAuth(key, sid string) string {
	val := url.Values{
		"client_secret": {key},
		"client_id":     {sid},
		"grant_type":    {"client_credentials"},
		"scope":         {"http://api.microsofttranslator.com"},
	}
	return val.Encode()
}

// Translator
type MicrosoftTranslator struct {
	Base
	auth
}
type auth struct {
	URL string
	VAL string
	tkn string
	tat time.Time
	ttl time.Duration
}

// Result
type TargetResult Result
type middleResult []byte

// KEY & SID
// Register your applications
// https://datamarket.azure.com/developer/applications/

// Token (POST)
// Content-Type: application/x-www-form-urlencoded
// https://datamarket.accesscontrol.windows.net/v2/OAuth2-13
// client_secret=key&client_id=sid&grant_type=client_credentials&scope=http://api.microsofttranslator.com
//
// {
//   "access_token": "http%3a%2f%2fschemas.xmlsoap.org%2fws%2f2005%2f05%2fidentity%2fclaims%2fnameidentifier=sid&http%3a%2f%2fschemas.microsoft.com%2faccesscontrolservice%2f2010%2f07%2fclaims%2fidentityprovider=https%3a%2f%2fdatamarket.accesscontrol.windows.net%2f&Audience=http%3a%2f%2fapi.microsofttranslator.com&ExpiresOn=1545329540&Issuer=https%3a%2f%2fdatamarket.accesscontrol.windows.net%2f&HMACSHA256=HafKUeEnoxntAWz4tWc6k10ZchCT-KC5mrD45tkMT7U",
//   "expires_in": "600",
//   "token_type": "http://schemas.xmlsoap.org/ws/2009/11/swt-token-profile-1.0",
//   "scope": "http://api.microsofttranslator.com"
// }

// Translate (GET)
// Authorization: Bearer http%3a%2f%2fschemas.xmlsoap.org%2fws%2f2005%2f05%2fidentity%2fclaims%2fnameidentifier=sid&http%3a%2f%2fschemas.microsoft.com%2faccesscontrolservice%2f2010%2f07%2fclaims%2fidentityprovider=https%3a%2f%2fdatamarket.accesscontrol.windows.net%2f&Audience=http%3a%2f%2fapi.microsofttranslator.com&ExpiresOn=1545329540&Issuer=https%3a%2f%2fdatamarket.accesscontrol.windows.net%2f&HMACSHA256=HafKUeEnoxntAWz4tWc6k10ZchCT-KC5mrD45tkMT7U
// http://api.microsofttranslator.com/v2/Http.svc/Translate?text=When%20you%20are%20old%20and%20grey%20and%20full%20of%20sleep,%20and%20nodding%20by%20the%20fire,%20take%20down%20this%20book,%20and%20slowly%20read.&to=zh&from=en
//
// <string xmlns="http://schemas.microsoft.com/2003/10/Serialization/">当你老了, 灰蒙蒙的, 睡着觉, 在火边点头的时候, 把这本书拿下来, 慢慢地读。</string>
