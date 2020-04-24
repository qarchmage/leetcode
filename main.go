package main

import "fmt"

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}))
}

func uniqueOccurrences(arr []int) bool {
	var count = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	var result = make(map[int]int)
	for _, value := range count {
		result[value]++
		if result[value] != 1 {
			return false
		}
	}
	return true
}
