package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
func lastStoneWeight(stones []int) int {
	var a, b, res int

	// smash until only one stone is left
	for len(stones) > 1 {

		sort.Ints(stones)
		//take two biggest
		a, b = stones[len(stones)-1], stones[len(stones)-2]
		//cut them from array
		stones = stones[:len(stones)-2]
		// smash them if  0 dont add
		res = abs(b - a)
		if res > 0 {
			stones = append(stones, res)
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}
	return 0

}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
