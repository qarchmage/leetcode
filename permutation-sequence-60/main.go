package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
	fmt.Println(getPermutation(3, 5))
}

// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}
func getPermutationNotRightOrderOnLeetCode(n int, k int) string {

	MATTE := ""

	for i := 1; i < n+1; i++ {
		MATTE = MATTE + strconv.Itoa(i)
	}

	result := []string{}
	Perm([]rune(MATTE), func(a []rune) {
		result = append(result, string(a))
	})
	return result[k-1]

}

func getPermutation(n int, k int) string {
	// offset k
	k = k - 1

	// build stack of permutations and an array from 1 to n
	permStack := []int{}
	nums := []int{1}
	currentP := 1
	for i := 2; i <= n; i++ {
		nums = append(nums, i)
		permStack = append(permStack, currentP)
		currentP = currentP * i
	}

	// build result
	result := ""
	for len(nums) > 1 {
		// pop from stack
		perm := permStack[len(permStack)-1]
		permStack = permStack[:len(permStack)-1]

		// find index
		idx := k / perm

		// find next k
		k = k % perm

		// add num at index to result
		result = result + strconv.Itoa(nums[idx])

		// delete num at index
		copy(nums[idx:], nums[idx+1:])
		nums = nums[:len(nums)-1]
	}

	return result + strconv.Itoa(nums[0])
}
