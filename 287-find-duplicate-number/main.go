package main

import "fmt"

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 3, 4}))
	fmt.Println(findDuplicate([]int{2, 2, 2, 2}))

	longSlice := []int{75, 75, 75, 75, 75, 91, 75, 75, 75, 75, 75, 75, 30, 75, 75, 78, 75, 75, 75, 75, 75, 7, 79, 93, 75, 75, 15, 75, 75, 75, 75, 75, 75, 4, 75, 75, 21, 75, 75, 19, 75, 59, 75, 76, 75, 75, 75, 75, 75, 75, 75, 33, 75, 75, 75, 58, 75, 75, 5, 75, 97, 81, 48, 42, 75, 87, 75, 75, 25, 27, 94, 20, 75, 75, 75, 29, 75, 75, 86, 67, 75, 75, 75, 65, 75, 75, 75, 75, 75, 39, 75, 56, 75, 75, 75, 75, 3, 75, 75, 75}

	fmt.Println(findDuplicate(longSlice))
}

func findDuplicate(nums []int) int {
	var x uint64
	max := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	max = (max / 64) + 1
	H := make([]uint64, max)
	for i := range nums {
		word, bit := nums[i]/64, nums[i]%64
		x = 1
		x = x << bit
		if x&H[word] > 0 {
			return nums[i]
		}
		H[word] |= x
	}
	return -1
}
