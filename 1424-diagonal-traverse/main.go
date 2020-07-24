package main

import (
	"fmt"
)

func main() {

	nums := [][]int{{5, 6, 3, 10}, {9}, {1, 19}, {9, 9, 2}}
	//nums := [][]int{{4, 8, 2, 5, 4}, {2, 3, 4, 2, 1}, {10, 9, 8, 7, 6}, {12, 19, 18, 4, 3}, {5, 4, 3, 2, 1}}
	fmt.Println(findDiagonalOrder(nums))

}

func findDiagonalOrder(nums [][]int) []int {
	m := make(map[int][]int)
	r := make([]int, 0)
	var s int
	// put each diagonal value with sum of i + j into map
	for i := range nums {
		for j := range nums[i] {
			m[i+j] = append(m[i+j], nums[i][j])
		}
	}
	for k := 0; k < len(m); k++ {
		//take last element first
		s = len(m[k])
		for j := s - 1; j >= 0; j-- {
			r = append(r, m[k][j])

		}
	}
	return r
}
