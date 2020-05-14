package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := generateSlice(5)
	sortColors(a)
	fmt.Println(a)
	fmt.Println(rand.Int() % 6)
	fmt.Println(rand.Int() % 6)
	fmt.Println(rand.Int() % 6)
}

func sortColors(nums []int) {
	quicksort(nums)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
