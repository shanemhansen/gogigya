package gogigya

import (
	"bytes"
	"crypto/hmac"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
    "crypto/sha1"
    "encoding/base64"
)


func New(this *Request) *Request {
    if this.Params == nil {
        this.Params = make(url.Values)
    }
	return this
}

type Request struct {
	ApiKey    string
	SecretKey string
	Method    string
	Format    string
	path      string
	domain      string
	Params    url.Values
}

func (this *Request) Send() (r *http.Response, err error) {
	if len(this.Format) == 0 {
		this.Format = "json"
	}
	this.Params.Add("format", "json")
	this.Params.Add("httpStatusCodes", "true")
	this.Params.Add("sdk", "php_2.15")
	this.Params.Add("apiKey", this.ApiKey)
	this.Params.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	this.Params.Add("nonce", fmt.Sprintf("%d", time.Now().UnixNano()/1000/1000))
	fields := strings.Split(this.Method, ".")
	this.domain = fields[0] + "." + "gigya.com"
	this.path = "/" + this.Method
    uri := "http://" + this.domain + this.path
    this.Params.Add("sig", GetSig(this.SecretKey, "POST", uri, false, this.Params))
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, vals := range this.Params {
		for _, val := range vals {
			err = writer.WriteField(key, val)
			if err != nil {
                return r, err
			}
		}
	}
	writer.Close()

	r, err = http.Post(
		"http://"+this.domain+this.path,
		"multipart/form-data; boundary="+ writer.Boundary(),
		body)
	return
}

type Response struct {
	ErrorCode int
}

func GetSig(secret, method, uri string, useHttps bool, params url.Values) string {
	base := CalcOauth1BaseString(method, uri, useHttps, params)
	return CalcSig(base, secret)
}
func CalcOauth1BaseString(method, uri string, secure bool, params url.Values) string {
	normalized := ""
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	protocol := u.Scheme
    ///ignore our problems we're always on the right port
	normalized += protocol + "://"
    normalized += u.Host
	normalized += u.Path
	keys := make([]string, 0)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	qs := ""
	amp := ""
	for _, k := range keys {
		for _, v := range params[k] {
			if strings.EqualFold(v, "0") {
				v = ""
			}
			qs += amp + k + "=" + Gigencode(v)
		}
		amp = "&"
	}
    qs = strings.Replace(qs, "+", "%20", -1)
	base := strings.Join([]string{strings.ToUpper(method),
		Gigencode(normalized),
		Gigencode(qs)}, "&")
	return base
}
func CalcSig(value, secret string) string {
	//value must be ut8
    key, err := base64.StdEncoding.DecodeString(secret)
    if err != nil {
        panic(err)
    }
    mac := hmac.New(sha1.New, key)
    mac.Write([]byte(value))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
    return sig
}
func Gigencode(value string) string {
	return strings.Replace(url.QueryEscape(value), "%7E", "~", -1)
}
