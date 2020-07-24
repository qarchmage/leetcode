package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("ab", "dbao"))
}

func checkInclusion(s1 string, s2 string) bool {
	uniqChars := make([]int, 26)

	for i := range s1 {
		uniqChars[s1[i]-'a']++
	}

	currSet := make([]int, 26)

	for i := range s2 {

		currSet[s2[i]-'a']++

		if i >= len(s1) {
			currSet[s2[i-len(s1)]-'a']--
		}
		//fmt.Println(currSet, uniqChars)
		if deepEqual(currSet, uniqChars) {
			return true
		}
	}

	return false
}

func deepEqual(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
