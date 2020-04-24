package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 4, 3, 2}
	fmt.Println(arrayPairSum(slice))
}

func arrayPairSum(nums []int) int {
	r := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		r += nums[i]
	}

	return r

}
