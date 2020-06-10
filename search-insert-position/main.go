package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

}
func searchInsert(nums []int, target int) int {
	// find index where to insert target in sorted array
	if target == 0 {
		return 0
	}

	for i := range nums {
		if target <= nums[i] {
			return i
		}
	}
	return len(nums)
}
