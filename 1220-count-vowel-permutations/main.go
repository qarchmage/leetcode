package main

import (
	"fmt"
)

func main() {
	fmt.Println(countVowelPermutation(144))
}

func countVowelPermutation(n int) int {
	a, e, i, o, u := 1, 1, 1, 1, 1
	for x := 0; x < n-1; x++ {
		a, e, i, o, u = e+i+u, a+i, e+o, i, i+o
	}
	return (a + e + i + o + u) % (1e9 + 7)
}
