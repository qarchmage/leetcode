package main

import "fmt"

func main() {
	citations := []int{0, 1, 3, 5, 6}
	c2 := []int{0, 1, 3, 5}
	fmt.Println(hIndex(citations))
	fmt.Println(hIndex(c2))
}

func hIndex(citations []int) int {
	n := len(citations)

	low := 0
	high := n - 1

	for low <= high {
		mid := low + (high-low)/2
		if citations[mid] == n-mid {
			return n - mid
		} else if citations[mid] > n-mid {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return n - low
}
