package main

import "fmt"

func main() {
	fmt.Println(generateTheString(105))

}

func generateTheString(n int) string {
	r := ""
	switch {
	case n == 1:
		return "a"
	case n%2 == 0:
		for i := 0; i < n-1; i++ {
			r += string(rune(97))
		}
		r += string(rune(98))
	default:
		for i := 0; i < n-2; i++ {
			r += string(rune(97))
		}
		r += string(rune(98))
		r += string(rune(99))
	}
	return r
}
