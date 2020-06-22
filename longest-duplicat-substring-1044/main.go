package main

import "fmt"

func main() {

	testcases := []string{"banana", "warcraft", "blub", "schlurp", "abcd", "starstar", "judas", "joseph", "glurfgl"}
	for _, v := range testcases {
		fmt.Printf("%q\n", longestDupSubstring(v))
	}
}

func longestDupSubstring(S string) string {
	length := len(S)

	power := make([]int, length)
	power[0] = 1

	for i := 1; i < length; i++ {
		power[i] = (power[i-1] * 26) //% prime
	}
	fmt.Println(power)

	l, r := 0, length-1
	out := ""

	for l <= r {
		mid := (l + r) / 2
		if s := RabinCorp(S, mid, power); len(s) != 0 {
			l = mid + 1
			out = s
		} else {
			r = mid - 1
		}
	}

	return out
}

func RabinCorp(s string, end int, power []int) string {

	if end == 0 {
		return ""
	}

	hash := map[int]bool{}
	currentHash := calculateHash(s, end)
	hash[currentHash] = true
	// fmt.Println("end", end, "currentHash", currentHash)

	for i := end; i < len(s); i++ {
		// fmt.Println("i", i)
		currentHash = (currentHash - power[end-1]*int(s[i-end])) // % prime
		// fmt.Println("currentHash", currentHash)
		currentHash = (currentHash*26 + int(s[i])) //% prime
		// fmt.Println("currentHash", currentHash)
		if _, ok := hash[currentHash]; ok {
			o := string(s[i-end+1 : i+1])
			return o
		} else {
			hash[currentHash] = true
		}
	}

	return ""
}

func calculateHash(s string, length int) int {
	h := 0
	for j := 0; j < length; j++ {
		h = (26*h + int(s[j])) // % prime
	}
	return h
}
