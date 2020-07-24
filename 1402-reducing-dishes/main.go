package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
}

func maxSatisfaction(satisfaction []int) int {
	quicksort(satisfaction)

	return findMaxProduct(satisfaction)
}

func findMaxProduct(s []int) int {
	var ans, cur, sum int
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		cur += sum + s[i]
		sum += s[i]

		ans = max(ans, cur)
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
