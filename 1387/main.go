package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getKth(1, 1000, 777))
}

//Step1: Find power of each number between low and high inclusive.
//Step2: Sort them by powers, if equal sort them by number itself.
//Step3; Return the kth number
type elem struct {
	num int
	pow int
}
type kpow []elem

func (e kpow) Len() int { return len(e) }
func (e kpow) Less(i, j int) bool {
	if e[i].pow < e[j].pow {
		return true
	}
	if e[i].pow > e[j].pow {
		return false
	}
	if e[i].num < e[j].num {
		return true
	}
	return false
}

func (e kpow) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func getKth(lo int, hi int, k int) int {
	powerK := make([]elem, hi-lo+1)
	for i := lo; i <= hi; i++ {
		powerK[i-lo] = elem{i, pow(i)}
	}
	sort.Sort(kpow(powerK))
	return powerK[k-1].num
}

func pow(num int) int {
	steps := 0
	for num > 1 {
		if num%2 == 0 { //even
			num = num / 2
		} else {
			num = num*3 + 1
		}
		steps++
	}
	return steps
}

// first try very slow
func getKthfirst(lo int, hi int, k int) int {
	if lo == 1 && hi == 1 && k == 1 {
		return 1
	}
	r := make(map[int]int)
	a := make([]int, hi-lo+1)
	j := 0
	min := 1001
	key := 1001
	for i := lo; i < hi+1; i++ {
		a[j] = i
		j++
	}

	for _, v := range a {
		x := v
		i := 0
		for {
			if x%2 == 0 {
				x /= 2
			} else {
				x = 3*x + 1
			}
			i++
			if x == 1 {
				r[v] = i
				break
			}
		}
	}
	// find min val starting point
	for i := 0; i < len(a); i++ {
		min = 1001
		key = 1001
		for j := lo; j < hi+1; j++ {
			if min > r[j] {
				min = r[j]
				key = j
			}
		}
		//fmt.Printf("k: %v min: %v\n", key, min)
		a[i] = key
		// mark as seen
		r[key] = 1001
	}

	return a[k-1]

}
