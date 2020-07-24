package main

import "fmt"

func main() {
	for i := 0; i < 31; i++ {

		fmt.Println(fib(i))
	}
}
func fib(N int) int {
	if N == 0 {
		return 0
	}
	// linear iterative
	// u can use recusion with memoization
	x, y := 0, 1
	for i := 0; i < N; i++ {
		x, y = x+y, x
	}

	return x

}
