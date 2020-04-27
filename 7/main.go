package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
}

func reverse(x int) int {
	tmp := 0
	r := ""
	isNegative, x := neg(x)
	for {
		tmp = x % 10
		r += strconv.Itoa(tmp)
		x /= 10
		if x == 0 {
			break
		}
	}

	if isNegative {
		x, _ = strconv.Atoi(r)
		x *= -1
	} else {
		x, _ = strconv.Atoi(r)
	}
	if math.MaxInt32 < x || x < math.MinInt32 {
		return 0
	}

	return x

}

func neg(x int) (bool, int) {
	if x > 0 {
		return false, x
	}
	x *= -1
	return true, x
}

func reverse2(x int) int {
	var reversed int
	for ; x != 0; x /= 10 {
		reversed *= 10
		reversed += x - ((x / 10) * 10)
	}
	if math.MaxInt32 < reversed || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}
