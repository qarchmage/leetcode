package main

import "fmt"

func main() {
	fmt.Println(singleNumberFast([]int{2, 2, 1, 3, 3, 5, 5, 6, 6}))
	a := 2
	a ^= 2
	fmt.Printf("%b", a)
}
func singleNumber(nums []int) int {
	m := make(map[int]int)
	ans := 0
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v == 1 {
			ans = k
		}
	}

	return ans
}

func singleNumberFast(nums []int) int {
	result := 0
	for _, el := range nums {
		fmt.Printf("result: %5d el: %5d\n", result, el)
		result ^= el
	}
	return result
}
