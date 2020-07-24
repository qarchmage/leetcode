package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{3, 1, 2, 4}))
}

func sortArrayByParity(A []int) []int {
	// even
	r := make([]int, 0, len(A))
	for i := range A {
		if A[i]%2 == 0 {
			r = append(r, A[i])
		}
	}

	// odds
	for i := range A {
		if A[i]%2 != 0 {
			r = append(r, A[i])
		}
	}
	//return
	return r

}
