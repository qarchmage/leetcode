package main

import "fmt"

func main() {
	a := []int{0, 1, 0, 3, 12, 0, 22, 33, 44, 0}
	moveZeroes(a)
	fmt.Println(a)
}
func moveZeroes(nums []int) {
	currentNotNullIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[currentNotNullIndex] = nums[i]

			if currentNotNullIndex != i {
				nums[i] = 0
			}

			currentNotNullIndex++
		}
	}
}
