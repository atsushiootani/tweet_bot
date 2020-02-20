package oauth1

/**
Auth1.0 を用いたHTTPリクエストを行う
 */

import (
	"../../http_util"
	"log"
	"net/http"
	"time"
)

// Auth1.0 用のリクエスト
type Request struct{
	http_util.Request
	credentials Credentials
}

func (req *Request) SetCredentials(credentials Credentials) {
	req.credentials = credentials
}

// Auth1.0 用のHttpリクエストを生成する
// interface HTTPRequester
func (req *Request) CreateHttpRequest() (hreq *http.Request, err error) {
	hreq, err = req.Request.CreateHttpRequest()
	if err != nil {
		return
	}

	hreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	hreq.Header.Add("Authorization", req.createAuthorizationHeaderString())

	return
}

// リクエストの実行
// interface HTTPRequester
func (req *Request) Do() (result string) {
	hreq, err := req.CreateHttpRequest()
	if err != nil{
		return
	}

	client := req.CreateClient(time.Duration(10) * time.Second)

	// call
	res, err := client.Do(hreq)
	if err != nil{
		log.Println("err client.Do", err)
		return
	}
	defer res.Body.Close()

	result = req.ResponseBodyAsString(res)
	return
}
