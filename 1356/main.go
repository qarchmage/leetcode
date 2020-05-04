package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		n, k := bits.OnesCount(uint(a)), bits.OnesCount(uint(b))
		if n != k {
			return n < k
		}
		return a < b
	})

	return arr
}
