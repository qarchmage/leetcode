package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(singleNumber([]int{2, 2, 2, 3}))
	fmt.Println(time.Since(start))

}

func singleNumber(nums []int) int {
	ones := uint(0)
	twos := uint(0)
	for i := uint(0); i < uint(len(nums)); i++ {
		ones = (ones ^ uint(nums[i])) &^ twos
		twos = (twos ^ uint(nums[i])) &^ ones
	}
	return int(ones)
}
