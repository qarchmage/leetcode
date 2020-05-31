package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "i love eating burger"
	searchWord := "burg"
	fmt.Println(isPrefixOfWord(sentence, searchWord))
}
func isPrefixOfWord(sentence string, searchWord string) int {
	a := strings.Split(sentence, " ")
	for i, v := range a {
		if strings.HasPrefix(v, searchWord) {
			return i + 1
		}
	}
	return -1

}
