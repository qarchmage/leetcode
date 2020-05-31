package main

import "fmt"

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}

func shortestToChar(S string, C byte) []int {
	n := len(S)

	res := make([]int, n)
	for i := range res {
		res[i] = n
	}

	l, r := -n, 2*n

	for i := 0; i < n; i++ {
		j := n - i - 1
		if S[i] == C {
			l = i
		}

		if S[j] == C {
			r = j
		}

		res[i] = min(res[i], dist(i, l))
		res[j] = min(res[j], dist(j, r))
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func dist(i, j int) int {
	if i < j {
		return j - i
	}

	return i - j
}
