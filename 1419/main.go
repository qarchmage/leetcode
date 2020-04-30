package main

import (
	"fmt"
)

func main() {
	testcases := []string{"croak", "croakcroak", "crcoakroak", "croakcroa", "croakcrook", "aoocrrackk"} // 1 1 2 -1 -1 -1
	for _, v := range testcases {
		fmt.Println(minNumberOfFrogs(v))
	}
}
func minNumberOfFrogs(croakOfFrogs string) int {
	c, r, o, a, k, frog := 0, 0, 0, 0, 0, 0
	for _, v := range croakOfFrogs {
		switch v {
		case 'c':
			c++
			if c > frog {
				frog = c
			}
		case 'r':
			r++
			if r > c {
				return -1
			}
		case 'o':
			o++
			if o > r {
				return -1
			}
		case 'a':
			a++
			if a > o {
				return -1
			}
		case 'k':
			k++
			if k > a {
				return -1
			}
			c--
			r--
			o--
			a--
			k--
		}
	}
	if c+r+o+a+k != 0 {
		return -1
	}
	return frog
}
