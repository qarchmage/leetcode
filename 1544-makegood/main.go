package main

import "fmt"

func main() {
	testcases := []string{"leeeEtcode", "abBAcC", "pachebel", "fFeinsSstein", "ddDeaAnfFnis ritgGchie", "cCkaAdDefFkKn"}
	for _, v := range testcases {
		fmt.Println(makeGood(v))
	}
}

func makeGood(s string) string {
	stack := []rune{}
	for _, r := range s {
		length := len(stack)
		last := len(stack) - 1
		if length > 0 && (stack[last] == r+32 || stack[last] == r-32) {
			stack = stack[:last]
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
