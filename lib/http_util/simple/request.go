package simple

/**
通常の HTTP リクエストを行う
 */

import (
	"../../http_util"
	"log"
	"net/http"
	"time"
)

type Request struct{
	http_util.Request
}

// Auth1.0 用のHttpリクエストを生成する
// interface HTTPRequester
func (req Request) CreateHttpRequest() (hreq *http.Request, err error) {
	return req.Request.CreateHttpRequest()
}

// Httpリクエストを実行する
// interface HTTPRequester
func (req Request) Do() (result string){
	hreq, err := req.CreateHttpRequest()
	if err != nil{
		return
	}

	client := req.CreateClient(time.Duration(10) * time.Second)

	// call
	res, err := client.Do(hreq)
	if err != nil{
		log.Print("err client.Do", err)
		return
	}
	defer res.Body.Close()

	result = req.ResponseBodyAsString(res)
	return
}
