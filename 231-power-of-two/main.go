package main

import "fmt"

func main() {
	a := []int{1, 16, 218}
	for _, v := range a {
		fmt.Println(isPowerOfTwo(v))
	}
}

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n > 1 {
		if n%2 == 1 {
			return false
		}

		n /= 2
	}

	return true
}
