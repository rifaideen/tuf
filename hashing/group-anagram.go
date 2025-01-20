package main

import (
	"sort"
	"strings"
)

func GroupAnagram(strs []string) [][]string {
	strMap := map[string][]string{}

	sortedKey := func(str string) string {
		key := sort.StringSlice(strings.Split(str, ""))
		key.Sort()

		return strings.Join(key, "")
	}

	for _, str := range strs {
		key := sortedKey(str)

		if _, ok := strMap[key]; !ok {
			strMap[key] = []string{}
		}

		strMap[key] = append(strMap[key], str)
	}

	ans := [][]string{}

	for _, val := range strMap {
		ans = append(ans, val)
	}

	return ans
}
