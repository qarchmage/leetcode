package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 200; i++ {
		a := ge()
		fmt.Println(maxSubarraySumCircular(a))
	}
	fmt.Println(time.Since(start))
}

func maxSubarraySumCircular(A []int) int {
	acc := 0
	max1 := kadane(A)

	for i := 0; i < len(A); i++ {
		acc += A[i]
		A[i] = -A[i]
	}

	min := kadane(A)
	max2 := acc + min
	if max2 == 0 {
		return max1
	}

	return int(math.Max(float64(max1), float64(max2)))
}

func kadane(a []int) int {
	cur := a[0]
	max := a[0]

	for i := 1; i < len(a); i++ {
		cur = int(math.Max(float64(cur+a[i]), float64(a[i])))
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max

}

func ge() []int {
	a := make([]int, rand.Intn(30000))
	min := -30000
	max := 30000
	for i := range a {
		a[i] = rand.Intn(max-min) + min
	}
	return a
}
