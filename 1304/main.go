package main

import "fmt"

func main() {
	fmt.Println(sumZero(5))
}

func sumZero(n int) []int {
	val := make([]int, n)
	for i := 0; i < n-1; i++ {
		val[i] = i + 1
	}
	fmt.Println(val)
	//n(n+1) / 2
	n--
	fmt.Println(n)
	val[n] = -(n * (n + 1)) / 2
	return val
}
