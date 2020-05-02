package main

import "fmt"

func main() {
	cases := make([]int, 25)
	for i := range cases {
		cases[i] = i
	}

	for _, v := range cases {
		fmt.Println(countLargestGroup(v))
	}
}

func countLargestGroup(n int) int {
	y := n
	cnt := 0
	for y > 0 {
		x := y % 10
		cnt++
		y -= x
		y /= 10
	}
	cache := make([]int, cnt*9+1)

	for i := 1; i <= n; i++ {
		y = i
		digits := 0
		for y > 0 {
			x := y % 10
			digits += x
			y -= x
			y /= 10
		}
		cache[digits]++
	}

	//find largest
	max := 0
	for i := 0; i < len(cache); i++ {
		if cache[i] > max {
			max = cache[i]
		}
	}

	ans := 0
	for i := 0; i < len(cache); i++ {
		if cache[i] == max {
			ans++
		}
	}

	return ans
}
