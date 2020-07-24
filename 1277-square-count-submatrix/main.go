package main

import (
	"fmt"
	"math"
)

func main() {
	a := [][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}}
	fmt.Println(countSquares(a))
}

func countSquares(matrix [][]int) int {
	dp := make([][]int, len(matrix)+1)
	for i := range dp {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	for i := range matrix {
		for j, v := range matrix[i] {
			if v == 0 {
				continue
			}
			tmp := min(dp[i][j+1], dp[i+1][j])
			if matrix[i-tmp][j-tmp] != 0 {
				dp[i+1][j+1] = tmp + 1
			} else {
				dp[i+1][j+1] = tmp
			}
		}
	}
	squares := 0
	for i := range dp {
		for _, v := range dp[i] {
			squares += v
		}
	}
	return squares
}

func min(values ...int) int {
	minValue := math.MaxInt32
	for _, v := range values {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}
