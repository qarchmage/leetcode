package main

import "fmt"

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}

func canBeEqual(target, arr []int) bool {
	// find max value
	max := 0
	for i := range target {
		if target[i] > max {
			max = target[i]
		}
	}
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	hash := make([]int, max+1)
	for i := range target {
		hash[target[i]]++
	}

	for i := range arr {
		hash[arr[i]]--
	}
	for i := range hash {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
