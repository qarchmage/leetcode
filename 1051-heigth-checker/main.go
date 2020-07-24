package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(input))
}
func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	copy(sorted, heights)
	sort.Ints(sorted)
	r := 0
	for i := range sorted {
		if sorted[i] != heights[i] {
			r++
		}
	}

	return r

}
