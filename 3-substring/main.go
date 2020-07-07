package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abc"))
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	maxLen, start := 0, 0
	table := [128]int{}
	for i := range table {
		table[i] = -1
	}
	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		maxLen = maxInt(maxLen, i-start+1)
	}
	return maxLen
}
