package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words))

}

func stringMatching(words []string) []string {

	r := make([]string, 0)

	for _, v := range words {
		for i := 0; i < len(words); i++ {
			if v == words[i] {
				continue
			}
			// if word is substring of other word append to result
			if strings.Contains(words[i], v) {
				r = append(r, v)
				break
			}
		}
	}
	return r

}
