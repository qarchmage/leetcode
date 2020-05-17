package main

import (
	"fmt"
)

func main() {
	cases := []string{"abab", "ab", "cbaebabacd", "abc", "", ""}
	for i := 0; i < len(cases); i += 2 {
		fmt.Println(findAnagrams(cases[i], cases[i+1]))
	}
}
func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var pat, mem [26]int
	// filling hashmaps
	for i := range p {
		pat[p[i]-'a']++
		mem[s[i]-'a']++
	}
	var out []int
	for i := 0; i < len(s)-len(p)+1; i++ {
		if pat == mem {
			out = append(out, i)
		}
		// sliding window, no need to change at last iteration
		if i+len(p) < len(s) {
			mem[s[i]-'a']--
			mem[s[i+len(p)]-'a']++
		}
	}
	return out
}

func findAnagrams(s string, p string) []int {
	r := []int{}
	if len(s) == 0 || len(p) == 0 {
		return r
	}
	search := []rune(s)
	initialIndex := 0
	searchChars := len(p)

	// init map
	m := make(map[rune]bool)
	for _, v := range p {
		m[v] = false
	}
	for {
		for i := initialIndex; i < initialIndex+searchChars; i++ {
			if isInMap(m, search[i]) {
				m[search[i]] = true
			}
		}
		if isTrue(m) {
			r = append(r, initialIndex)
		}
		initialIndex++
		m = resetMap(m)
		if initialIndex > len(s)-searchChars {
			break
		}
	}
	return r
}

func isTrue(m map[rune]bool) bool {
	for _, v := range m {
		if v == false {
			return false
		}
	}
	return true
}

func isInMap(m map[rune]bool, r rune) bool {
	for k := range m {
		if k == r {
			return true
		}
	}
	return false
}

func resetMap(m map[rune]bool) map[rune]bool {
	for k := range m {
		m[k] = false
	}
	return m
}
