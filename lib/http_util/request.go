package http_util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HTTP リクエストを行うインタフェース
type HTTPRequester interface{
	CreateHttpRequest()
	Do()
}

// HTTP リクエスト
type Request struct{
	Method Method
	Url string
	Params Params
	Headers Params
}

// HTTP リクエストを作成
func (req *Request) CreateHttpRequest() (hreq *http.Request, err error) {

	// params
	values := url.Values{}
	for k, v := range req.Params {
		values.Add(k, v)
	}

	// hreq
	hreq, err = http.NewRequest(req.Method.ToString(), req.Url, strings.NewReader(values.Encode()))
	if err != nil{
		logger.Print("err NewRequest", err)
		return
	}

	// headers
	for k, v := range req.Headers {
		hreq.Header.Add(k, v)
	}

	req.log()
	return
}

// クライアントの作成
func (req *Request) CreateClient(timeout time.Duration) *http.Client{
	return &http.Client{Timeout: timeout}
}

func (req *Request) ResponseBodyAsJson(res *http.Response, data interface {}) error {
	body := req.ResponseBody(res)

	err := json.Unmarshal(body, &data)
	if err != nil {
		logger.Print("err ", err)
	}

	return err
}

func (req *Request) ResponseBodyAsString(res *http.Response) string{
	return string(req.ResponseBody(res))
}

func (req *Request) ResponseBody(res *http.Response) []byte{
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		logger.Print("err readBody", err)
		return nil
	}

	return body
}

func (req *Request) log() {
	s := fmt.Sprintf("http_util.Request: Method=%s, Url=%s, Params=%s, Headers=%s", req.Method, req.Url, req.Params, req.Headers)
	logger.Output(2, s)
}
