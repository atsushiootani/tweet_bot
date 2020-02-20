package http_util

import (
	"../util/map"
	"net/url"
	"strings"
)

type Params map[string]string

// クエリパラメータをカンマ区切りの文字列にして返す
//
// ex.
// key1="value1", key2="value2"
func (param Params) QueryParameterToCommaString() string{
	result := strings.Builder{}

	for k, v := range param {
		if result.Len() > 0 {
			result.WriteString(", ")
		}

		result.WriteString(k + "=\"" + v + "\"")
	}
	return result.String()
}

// クエリパラメータを文字列にして返す
func (param Params) QueryParameterToString() string{
	result := strings.Builder{}

	for k, v := range param {
		if result.Len() > 0 {
			result.WriteString("&")
		}
		result.WriteString(k + "=" + url.QueryEscape(v))
	}
	return result.String()
}

// クエリパラメータをキー名順に並べた文字列にして返す
func (param Params) QueryParameterToSortedString() string{
	result := strings.Builder{}

	sortedKeys := _map.SortedKeysOfMap(param)
	for _, k := range sortedKeys {
		v := param[k]

		if result.Len() > 0 {
			result.WriteString("&")
		}
		result.WriteString(k + "=" + url.QueryEscape(v))
	}
	return result.String()
}

func MergeParam(param1, param2 Params) Params{
	return _map.MergeMap(param1, param2)
}
