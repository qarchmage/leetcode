package main

import "fmt"

func main() {
	fmt.Println(findLucky([]int{5}))
}

func findLucky(arr []int) int {
	m := make(map[int]int)
	r := 0
	for _, v := range arr {
		m[v]++
	}
	for k, v := range m {
		if k == v {
			if v > r {
				r = v
			}
		}
	}

	if r == 0 {
		return -1
	}

	return r

}
