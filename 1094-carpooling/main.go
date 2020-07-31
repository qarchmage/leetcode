package main

import "fmt"

func main() {
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 4))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 3, 7}}, 5))
	fmt.Println(carPooling([][]int{{2, 1, 5}, {3, 5, 7}}, 3))
	fmt.Println(carPooling([][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, 11))
}

func carPooling(trips [][]int, capacity int) bool {
	stops := [1001]int{}

	for _, v := range trips {
		stops[v[1]] += v[0]
		stops[v[2]] -= v[0]
	}

	var usedSeats int

	for _, v := range stops {
		usedSeats += (v)
		if usedSeats > capacity {
			return false
		}
	}

	return true
}
