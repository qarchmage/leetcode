package main

import "fmt"

func main() {
	a := [21]int{}
	b := [21]bool{}
	for i := range a {
		a[i] = i
		b[i] = isHappy(i)
	}
	fmt.Println(a)
	fmt.Println(b)

}
func isHappy(n int) bool {
	x, y := n, n
	for {
		x = digitSum(x)
		y = digitSum(digitSum(y))
		if x == 1 || y == 1 {
			return true
		}
		if x == y {
			return false
		}
	}
}

func digitSum(n int) int {
	var digitSlice []int
	for n != 0 {
		digitSlice = append(digitSlice, n%10)
		n /= 10
	}

	var sum int
	for _, v := range digitSlice {
		sum += v * v
	}
	return sum
}
