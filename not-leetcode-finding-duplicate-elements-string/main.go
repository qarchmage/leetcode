package main

import "fmt"

func main() {
	findDuplicateCharInString("finding")
	var i uint32
	i = (1 << 32) - 1 // max value of uint32
	fmt.Println(i)
}

// print all duplicates in string (bitwise operation)
func findDuplicateCharInString(s string) {
	var H, x uint32
	for i := 0; i < len(s); i++ {
		x = 1
		x = x << (s[i] - 97)
		if x&H > 0 {
			fmt.Printf("%c is duplicate\n", s[i])
		} else {
			H = x | H
		}

	}

}
