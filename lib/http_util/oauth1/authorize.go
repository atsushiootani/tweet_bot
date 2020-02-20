package oauth1

/**
  auth1.0 に関する処理
 */

import (
	"../../http_util"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"
)

// データとキーから、署名文字列を返す
func createSignature(data, key string) string{
	str := hmacSha1(data, key)
	// base64String := base64.URLEncoding.EncodeToString([]byte(str))
	// return url.QueryEscape(base64String)
	return url.QueryEscape(str)
}

// HMAC-SHA1 変換を行う
func hmacSha1(data, key string) string{
	keyBytes := ([]byte)(key)
	hash := hmac.New(sha1.New, keyBytes)
	io.WriteString(hash, data)
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// 署名キーを作る
func createSignatureKey(consumerSecret, tokenSecret string) string {
	return fmt.Sprintf("%s&%s", url.QueryEscape(consumerSecret), url.QueryEscape(tokenSecret))
}

// 署名キーを作る
// リクエストトークン取得時は、シークレットがないのでこちらのメソッドを使うこと
func createSignatureKeyWithoutSecret(consumerSecret string) string{
	return createSignatureKey(consumerSecret, "")
}

// 署名対象のテキストを生成する
func createSignatureBaseString(method, api string, params http_util.Params, authorizationHeader http_util.Params) string{
	mergedParam := http_util.MergeParam(params, authorizationHeader)
	paramString := mergedParam.QueryParameterToSortedString()

	return fmt.Sprintf("%s&%s&%s", url.QueryEscape(method), url.QueryEscape(api), url.QueryEscape(paramString))
}

// 署名文字列を作成
func (req *Request) createSignatureString() string{
	signatureKey := createSignatureKey(req.credentials.ApiSecretKey, req.credentials.AccessTokenSecret)

	baseString := createSignatureBaseString(req.Request.Method.ToString(), req.Request.Url, req.Request.Params, req.createAuthHeaders())

	return createSignature(baseString, signatureKey)
}

// authHeader を生成
func (req *Request) createAuthHeaders() http_util.Params{
	return http_util.Params{
		"oauth_consumer_key": req.credentials.ApiKey,
		"oauth_nonce": getUnixtimeAsString(),
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp": getUnixtimeAsString(),
		"oauth_token": req.credentials.AccessToken,
		"oauth_version": "1.0",
	}
}

// Authorization ヘッダの値となる文字列を生成して返す
func (req *Request) createAuthorizationHeaderString() string{
	param := req.createAuthHeaders()
	param["oauth_signature"] = req.createSignatureString()

	return "OAuth " + param.QueryParameterToCommaString()
}

func getUnixtimeAsString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
