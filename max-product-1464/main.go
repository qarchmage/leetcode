package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
}

func maxProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return (nums[l-2] - 1) * (nums[l-1] - 1)
}
