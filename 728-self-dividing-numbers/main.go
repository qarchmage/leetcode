package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {
	numbers := make([]int, 0, right)
	result := make([]int, 0, right)
	n := 0
	for i := left; i < right+1; i++ {
		numbers = append(numbers, i)
	}
	for _, v := range numbers {
		n = v
		for n > 0 {
			tmp := n % 10
			if tmp == 0 || (v%tmp) != 0 {
				break
			}
			n /= 10

		}
		if n == 0 {
			result = append(result, v)
		}
	}

	return result

}
