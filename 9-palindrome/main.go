package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
}

func isPalindrome(x int) bool {

	it := strconv.Itoa(x)
	palindrome := make([]byte, 0)

	for i := len(it) - 1; i > -1; i-- {
		palindrome = append(palindrome, it[i])

	}

	if it == string(palindrome) {
		return true
	}

	return false
}
