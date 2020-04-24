package main

import "fmt"

func main() {
	fmt.Println(hammingDistance(1, 8))
}

func hammingDistance(x int, y int) int {

	diff := x ^ y
	bits := 0
	for diff > 0 {
		if diff&1 != 0 {
			bits++
		}
		diff = diff >> 1
	}
	return bits

}
