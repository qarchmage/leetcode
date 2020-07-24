package main

import "fmt"

func main() {
	s := "rat"
	fmt.Println(sortString(s))
}

func sortString(s string) string {
	m := [26]int{}

	for _, v := range s {
		m[v-'a']++
	}

	rtn := []byte{}
	for len(rtn) != len(s) {

		for i := 0; i < 26; i++ {
			if m[i] != 0 {
				rtn = append(rtn, byte('a'+i))
				m[i]--
			}
		}

		for i := 25; i >= 0; i-- {
			if m[i] != 0 {
				rtn = append(rtn, byte('a'+i))
				m[i]--
			}
		}
	}

	return string(rtn)

}
