package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))
}

func minimumAbsDifference(arr []int) [][]int {
	if len(arr) <= 1 {
		return nil
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		newdiff := arr[i] - arr[i-1]
		if diff > newdiff {
			diff = newdiff
		}
	}

	res := [][]int{}
	for i := 1; i < len(arr); i++ {
		newdiff := arr[i] - arr[i-1]
		if diff == newdiff {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}
