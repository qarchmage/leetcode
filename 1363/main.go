package main

import "fmt"

func main() {
	s := "011101"
	fmt.Println(maxScore(s))
}

func maxScore(s string) int {
	count0, count1 := 0, 0
	for _, v := range s {
		if v == '0' {
			count0++
		} else {
			count1++
		}
	}
	currentScore, maxScore := count1, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			currentScore++
		} else {
			currentScore--
		}
		if currentScore > maxScore {
			maxScore = currentScore
		}
	}
	return maxScore
}
