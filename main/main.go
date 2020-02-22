package main

import (
	"fmt"
	"sort"
	"strings"
)

func contains(array1[] string, str string) bool {
	for _, inStr := range array1 {
		if str == inStr {
			return true
		}
	}

	return false
}

func InArray(array1 []string, array2 []string) []string {
	var ans = make([]string, 0)

	for _, str1 := range array1 {
		if ! contains(ans, str1) {
			for _, str2 := range array2 {
				if strings.Contains(str2, str1) {
					ans = append(ans, str1)
					break
				}
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})

	return ans
}

func main() {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	//var r = []string{"arp", "live", "strong"}

	fmt.Println(InArray(a1, a2))
}
