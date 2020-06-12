package main

import (
	"fmt"
)

func main() {
	a := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	fmt.Println(diagonalSort(a))
}

func diagonalSort(mat [][]int) [][]int {
	y := len(mat)
	if y == 0 {
		return [][]int{}
	}
	x := len(mat[0])

	var nums [101]int

	for i := y; i >= 0; i-- {
		diagonal(mat, 0, i, &nums)
		replaceDiagonal(mat, 0, i, &nums)
	}

	for i := 1; i < x; i++ {
		diagonal(mat, i, 0, &nums)
		replaceDiagonal(mat, i, 0, &nums)
	}

	return mat
}

func diagonal(mat [][]int, x, y int, nums *[101]int) {
	for count := 0; y+count < len(mat) && x+count < len(mat[0]); count++ {
		nums[mat[y+count][x+count]]++
	}
}

func replaceDiagonal(mat [][]int, x, y int, nums *[101]int) {
	var count int
	for i := range nums {
		for nums[i] > 0 {
			mat[y+count][x+count] = i
			count++
			nums[i]--
		}
	}
}
