package main

import "fmt"

func main() {
	a := [][]int{{-1, 0}, {2, -1}}
	fmt.Println(countNegatives(a))
}
func countNegatives(grid [][]int) int {
	n := 0
	for _, slice := range grid {
		for _, number := range slice {
			if number < 0 {
				n++
			}
		}
	}
	return n

}
