package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
	fmt.Println(freqAlphabets(s))
}

func freqAlphabets(s string) string {
	r := ""
	slice := []rune(s)
	fmt.Println(slice)
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == '#' {
			r = string(96+((slice[i-1]-48)+(slice[i-2]-48)*10)) + r
			i -= 2
		} else {
			r = string(slice[i]+48) + r
		}
	}
	return r
}

func freqAlphabets2(s string) string {
	var (
		r, tmp string
		tmp2   int
	)
	slice := make([]rune, 0, len(s))
	for _, v := range s {
		slice = append(slice, v)
	}
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == '#' {
			tmp = string(slice[i-2]) + string(slice[i-1])
			tmp2, _ = strconv.Atoi(tmp)
			r = string(rune(tmp2+96)) + r
			i -= 2
		} else {
			r = string(slice[i]+48) + r
		}
	}
	return r

}
