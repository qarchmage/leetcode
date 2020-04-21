package main

import (
	"fmt"
)

func main() {
	a := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(a))
}

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	b := ""
	m := make(map[string]int)
	for _, word := range words {
		for _, rune := range word {
			rune -= 'a'
			b += morse[rune]
		}
		m[b]++
		b = ""
	}

	return len(m)
}
