package main

import "fmt"

func main() {

	nums := [][]int{{-3, 2, -3, 4, 2}, {1, 2}, {1, -2, -3}}

	for _, v := range nums {
		fmt.Println(minStartValue(v))
	}

}

func minStartValue(nums []int) int {
	var startValue, tmp int

	// loop until value is greater than 0
	for {
		// for each step increment startValue by 1
		startValue++
		// set to temp variable to not lose initial startValue which we want to return
		tmp = startValue
		for _, v := range nums {
			tmp += v
			// if tmp is less than 1 break loop and start over with startValue + 1
			if tmp <= 0 {
				break
			}
		}
		// check for return if tmp greater than 0 return
		if tmp > 0 {
			return startValue
		}
	}
}
