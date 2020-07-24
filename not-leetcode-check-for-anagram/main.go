package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "decimal"
	b := "medical"
	fmt.Println(isAnagram(a, b))
}

// find anagram linear time complexity O(n)
func isAnagram(a, b string) bool {
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	s := [26]int{}
	for i := range a {
		s[a[i]-97]++
	}

	for i := range b {
		s[b[i]-97]--
	}

	for i := range s {
		if s[i] != 0 {
			return false
		}
	}
	return true
}
