package main

import "fmt"

func main() {
	input := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(flipAndInvertImage(input))
	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 1)
}
func flipAndInvertImage(A [][]int) [][]int {
	revert := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		for j := len(A[i]) - 1; j > -1; j-- {
			revert[i] = append(revert[i], A[i][j])
		}
	}

	for i := 0; i < len(revert); i++ {
		for j := 0; j < len(revert[i]); j++ {
			revert[i][j] ^= 1
		}
	}

	return revert
}
