package main

import "fmt"

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))

}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for i := range candies {
		if candies[i] > max {
			max = candies[i]
		}
	}
	r := make([]bool, 0, len(candies))

	for i := range candies {
		if candies[i]+extraCandies >= max {
			r = append(r, true)
		} else {
			r = append(r, false)
		}
	}
	return r

}
