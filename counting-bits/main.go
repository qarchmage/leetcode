package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		fmt.Println(countBits(i))
	}
}

func countBits(num int) []int {
	f := make([]int, num+1)
	for i := 1; i <= num; i++ {
		f[i] = f[i>>1] + (i & 1)
	}
	return f
}
