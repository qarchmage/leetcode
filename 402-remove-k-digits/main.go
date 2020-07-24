package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	b := []int{3, 1, 2, 2, 1, 2, 2, 2, 3, 2, 2, 2, 2}
	a := []string{"1432219", "10200", "10", "1111", "112", "1155", "1123", "1234567890", "9876543210", "110", "3033", "001", "11220"}
	start := time.Now()
	for i := range a {
		fmt.Printf("input: %15s %5d \tresult: %15s\n", a[i], b[i], removeKdigits(a[i], b[i]))
	}

	elapsed := time.Since(start)
	log.Printf("first took %s", elapsed)
	fmt.Println("second")

	start = time.Now()
	for i := range a {
		fmt.Printf("input: %15s %5d \tresult: %15s\n", a[i], b[i], removeKdigits2(a[i], b[i]))
	}
	elapsed = time.Since(start)
	log.Printf("first took %s", elapsed)
}

// faster solution with stack
func removeKdigits2(num string, k int) string {
	if len(num)-k == 0 {
		return "0"
	}
	st := newStack(len(num))
	for _, n := range num {
		if st.empty() {
			st.push(n)
			continue
		}
		for !st.empty() && st.back() > int(n-'0') && k > 0 {
			st.pop()
			k--
		}
		st.push(n)
	}
	for !st.empty() && k > 0 {
		st.pop()
		k--
	}
	return st.String()
}

type stack struct {
	idx int
	s   []int
}

func newStack(size int) *stack {
	return &stack{
		s: make([]int, size),
	}
}

func (s *stack) push(num rune) {
	s.s[s.idx] = int(num - '0')
	s.idx++
}

func (s *stack) pop() {
	s.idx--
}

func (s *stack) back() int {
	return s.s[s.idx-1]
}

func (s *stack) empty() bool {
	return s.idx == 0
}

func (s *stack) String() string {
	var idx int
	for ; idx < s.idx; idx++ {
		if s.s[idx] == 0 {
			continue
		}
		break
	}
	if idx >= s.idx {
		return "0"
	}
	res := make([]rune, 0, len(s.s)-idx)
	for ; idx < s.idx; idx++ {
		res = append(res, rune(s.s[idx]+'0'))
	}
	return string(res)
}
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}
	for ; k > 0; k-- {
		length := len(num)
		i := 0
		for i+1 < length && num[i] <= num[i+1] {
			i += 1
		}
		num = num[:i] + num[i+1:]
	}

	// remove leading 0s
	length := len(num)
	for {
		// TrimPrefix only remove one 0
		num = strings.TrimPrefix(num, "0")
		if length == len(num) {
			break
		}
		length = len(num)
	}
	if len(num) == 0 {
		return "0"
	}
	return num
}

// NOO !!!!!!! not working

//func removeKdigits(num string, k int) string {
//r := []rune(num)
//if k == len(r) || k == len(r)-1 && r[len(r)-1] == '0' {
//return "0"
//}
//// check if last element is 0 ??
//storeindex := make([]int, 0, k)
//count := 0
//count, storeindex = findDig(r, len(r), storeindex, count, k)
//// if coutn is not equal to k reset and search reverse
//if count != k {
//count = 0
//storeindex = []int{}
//}
//// if count = 0 it is ascending digits only
//if count == 0 {
//count, storeindex = findDigReverse(r, len(r), storeindex, count, k)
//for i := 0; i < len(storeindex); i++ {
//r = del(r, storeindex[i])
//}
//} else {
//for i := len(storeindex) - 1; i >= 0; i-- {
//r = del(r, storeindex[i])
//}
//}
//// if count is still 0 means all numbers are same => remove k digits
//if count == 0 {
//r = r[:len(r)-k]
//return string(r)
//}
//for {
//// remove leading zero
//if r[0] == '0' && len(r) > 1 {
//r = del(r, 0)
//} else {
//break
//}
//}

//return string(r)

//}

//func del(a []rune, i int) []rune {
//if i == len(a) {
//a = a[:i-1]
//return a
//}
//copy(a[i:], a[i+1:])
//a[len(a)-1] = 0 // or the zero value of T
//a = a[:len(a)-1]
//return a
//}

//func findDig(r []rune, l int, s []int, c int, k int) (int, []int) {
//for i := 1; i < len(r); i++ {
//// if digit is less than the previous digit, discard the previous one
//if r[i] < r[i-1] {
//s = append(s, i-1)
//c++
//if c == k {
//break
//}
//}
//}
//return c, s
//}

//func findDigReverse(r []rune, l int, s []int, c int, k int) (int, []int) {
//for i := len(r) - 1; i > 0; i-- {
//// if digit is less than the previous digit, discard the previous one
//if r[i] > r[i-1] {
//s = append(s, i)
//c++
//if c == k {
//break
//}
//}
//}
//return c, s
//}
