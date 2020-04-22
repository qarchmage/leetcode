package main

import "fmt"

func main() {
	fmt.Println(repeatedNTimes([]int{4, 2, 3, 3}))
}

func repeatedNTimes(A []int) int {
	m := make(map[int]int)
	s := len(A) / 2
	r := 0
	for _, v := range A {
		m[v]++
	}
	for k, v := range m {
		if v == s {
			r = k
		}
	}

	return r

}
