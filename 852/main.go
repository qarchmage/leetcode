package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
}

func peakIndexInMountainArray(A []int) int {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return 0
}
