package main

import "fmt"

func main() {
	fmt.Println(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
	fmt.Println(busyStudent([]int{4}, []int{4}, 4))
	fmt.Println(busyStudent([]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7))
	fmt.Println(busyStudent([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5))
}

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var ans int
	for i, v := range startTime {
		if v <= queryTime {
			if endTime[i] >= queryTime {
				ans++
			}
		}
	}
	return ans
}
