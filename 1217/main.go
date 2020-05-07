package main

import "fmt"

func main() {
	fmt.Println(minCost([]int{2, 2, 2, 3, 3}))
}
func minCost(chips []int) int {
	odd, even := 0, 0
	for _, c := range chips {
		if c%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if even < odd {
		return even
	}

	return odd
}
