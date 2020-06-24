package main

import "fmt"

func main() {
	fmt.Println(numTrees(4))
}

// shouldn't it be called how many permutations ?
func numTrees(n int) int {
	if n <= 0 {
		return 0
	}

	// n+1 is enough, n+2 just prevent t[2] index out of range
	t := make([]int, n+2)
	t[0], t[1], t[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			if j >= i-j-1 {
				t[i] *= 2
				if j == i-j-1 {
					t[i] += t[j] * t[i-j-1]
				}
				break
			}
			t[i] += t[j] * t[i-j-1]
		}
	}

	fmt.Println(t)
	return t[n]
}
