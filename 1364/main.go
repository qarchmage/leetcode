package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
}

func minSubsequence(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	var ans []int
	totalSum := 0
	ansSum := 0
	sort.Ints(nums)
	fmt.Println(nums)

	for i := range nums {
		totalSum += nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		ansSum += nums[i]
		totalSum -= nums[i]
		ans = append(ans, nums[i])

		if ansSum > totalSum {
			break
		}
	}

	return ans

}
