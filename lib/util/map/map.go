package _map

import "sort"

// mapのキーをソートして返す
func SortedKeysOfMap(val map[string]string) []string {
	result := KeysOfMap(val)
	sort.Strings(result)
	return result
}

// mapのキーだけを返す
func KeysOfMap(val map[string]string) []string {
	var result []string

	for k, _ := range val{
		result = append(result, k)
	}

	return result
}

// map のマージ
func MergeMap(map1, map2 map[string]string) map[string]string {
	result := make(map[string]string)

	for k, v := range map1 {
		result[k] = v
	}

	for k, v:= range map2 {
		result[k] = v
	}
	return result
}

