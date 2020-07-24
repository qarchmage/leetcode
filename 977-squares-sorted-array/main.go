package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sortedSquares([]int{-2, 0}))
}

func sortedSquares(A []int) []int {

	for i := range A {
		A[i] *= A[i]
	}
	sort.Ints(A)

	return A
}
