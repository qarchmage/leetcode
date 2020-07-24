package main

import "fmt"

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 2, 2, 3, 3, 4, 4, 8, 8}))
}
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
