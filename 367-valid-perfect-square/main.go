package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPerfectSquare(81))
}

func isPerfectSquare(num int) bool {
	root := Sqrt(float64(num))
	if int(root+0.5)*int(root+0.5) == num {
		//is true
		return true
	} else {
		// is not
		return false
	}

}

func Sqrt(x float64) float64 {
	prev, z := float64(0), float64(1)
	for abs(prev-z) > 1e-8 {
		prev = z
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func abs(number float64) float64 {
	if number < 0 {
		return number * -1
	} else {
		return number
	}
}
