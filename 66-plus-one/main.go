package main

import (
	"fmt"
)

func main() {
	fmt.Println(plusOne([]int{9, 9}))

	fmt.Println(plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}

func plusOne(d []int) []int {
	c := 1

	for i := len(d) - 1; i >= 0; i-- {
		d[i], c = (d[i]+c)%10, (d[i]+c)/10
	}

	if c != 0 {
		return append([]int{c}, d...)
	}

	return d
}
