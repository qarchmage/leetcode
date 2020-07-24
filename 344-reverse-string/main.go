package main

import "fmt"

func main() {
	a := []string{"hello", "matthias", "LICH", "ff", "asdf"}
	for _, v := range a {
		b := []byte(v)
		reverseString(b)
		fmt.Println(string(b))
	}
}
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
