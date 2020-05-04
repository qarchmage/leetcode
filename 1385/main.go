package main

import (
	"fmt"
)

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ans := 0
	for i := 0; i < len(arr1); i++ {
		ans++
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i], arr2[j]) <= d {
				ans--
				break
			}
		}
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
