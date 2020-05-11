package main

import "fmt"

func main() {
	fmt.Println(longestPrefix("leetcodeleet"))
}

func longestPrefix(s string) string {
	return s[0:kmp(s)[len(s)-1]]
}

func kmp(s string) []int {
	ret := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := ret[i-1]
		for j > 0 && s[i] != s[j] {
			j = ret[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		ret[i] = j
	}
	return ret
}
