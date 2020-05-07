package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sumFourDivisors([]int{2, 21, 1}))

}

func sumFourDivisors(nums []int) int {
	r := 0
	for _, v := range nums {
		r += sumOfDivCount(v)
	}
	return r
}

func sumOfDivCount(num int) int {
	if num <= 5 {
		return 0
	}
	root := int(math.Sqrt(float64(num)))
	if root*root == num {
		return 0
	}
	divisorCount := 2
	summary := num + 1
	for i := 2; i <= root; i++ {
		if num%i == 0 {
			divisorCount += 2
			summary += i
			summary += num / i
			if divisorCount > 4 {
				return 0
			}
		}
	}
	if divisorCount == 4 {
		return summary
	}
	return 0
}
