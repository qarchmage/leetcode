package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(7, 3))
}

func uniquePaths(m int, n int) int {
	left := 1
	up := make([]int, m)

	for i := range up {
		up[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			up[j] = up[j] + left
			left = up[j]
		}
		left = 1
	}

	return up[m-1]
}
