package main

import "fmt"

func main() {
	a := [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}
	fmt.Println(luckyNumbers(a))
}

func luckyNumbers(matrix [][]int) []int {
	min := []int{}
	tmp := 0
	// find min values in row
	for i, slice := range matrix {
		tmp = matrix[i][0]
		for _, number := range slice {
			if tmp > number {
				tmp = number
			}
		}
		min = append(min, tmp)
	}
	//find max values in col

	f := 0
	maxf := 0
	mn := []int{}
max:
	for i := range matrix {
		if maxf < matrix[i][f] {
			maxf = matrix[i][f]
		}

	}
	mn = append(mn, maxf)
	if f < len(matrix[0])-1 {
		maxf = 0
		f++
		goto max
	}

	//find lucky
	for _, number := range min {
		for i := 0; i < len(mn); i++ {
			if number == mn[i] {
				return []int{number}
			}
		}
	}

	return nil

}
