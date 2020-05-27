package main

import "fmt"

func main() {
	a := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19, 19}
	b := []int{2, 1, 4, 3, 9, 6}
	// expected : 2,2,2,1,4,3,3,9,6,7,19
	fmt.Println(relativeSortArray(a, b))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	//find max number
	max := 0
	for i := range arr1 {
		if arr1[i] > max {
			max = arr1[i]
		}
	}
	//hash arr1 increment by 1 for each number
	hash := make([]int, max+1)
	for i := range arr1 {
		hash[arr1[i]]++
	}
	result := make([]int, 0)
	// first append number from arr2 then for each count of hash of that number
	for i := range arr2 {
		result = append(result, arr2[i])
		if hash[arr2[i]] > 0 {
			for j := 1; j < hash[arr2[i]]; j++ {
				result = append(result, arr2[i])
			}
			// if finshed set number count to zero
			hash[arr2[i]] = 0
		}
	}
	// add remaining numbers which were not in arr2
	// add from left to right (SORTED)
	for i := range hash {
		if hash[i] > 0 {
			// check if count is greater than 1
			for j := 0; j < hash[i]; j++ {
				result = append(result, i)
			}
		}

	}

	return result
}
