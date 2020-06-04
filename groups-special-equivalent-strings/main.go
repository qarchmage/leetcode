package main

import "fmt"

func main() {
	fmt.Println(numSpecialEquivGroups([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
}
func numSpecialEquivGroups(A []string) int {
	m := map[[52]int]interface{}{}
	for _, s := range A {
		fmt.Println(s)
		array := [52]int{}
		for i, c := range s {
			if i%2 == 0 {
				array[c-'a']++
			} else {
				array[c-'a'+26]++
			}
		}
		m[array] = nil
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	return len(m)
}
