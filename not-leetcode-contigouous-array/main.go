package main

import "fmt"

func main() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}

func findMaxLength(nums []int) int {
	for i := range nums {
		if nums[i] == 0 {
			nums[i] = -1
		}
	}
	sum, max := 0, 0
	m := make(map[int]int)
	m[0] = -1
	for i := range nums {
		sum += nums[i]
		if _, ok := m[sum]; ok {
			max = maxInt(max, i-m[sum])
		} else {
			m[sum] = i
		}
	}

	return max

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
