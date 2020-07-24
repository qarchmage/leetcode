package main

import "fmt"

func main() {
	n1 := 3
	n2 := 2
	n3 := 3
	n4 := 4
	r4 := [][]int{{4, 3}, {1, 4}, {4, 6}, {1, 7}}
	n5 := 4
	r5 := [][]int{{2, 10}, {3, 1}, {1, 2}, {2, 2}, {3, 5}, {4, 1}, {4, 9}, {2, 7}}
	r := [][]int{{1, 2}, {1, 3}, {1, 8}, {2, 6}, {3, 1}, {3, 10}}
	r2 := [][]int{{2, 1}, {1, 8}, {2, 6}}
	r3 := [][]int{{2, 3}}
	fmt.Println(maxNumberOfFamilies(n1, r))  // 4
	fmt.Println(maxNumberOfFamilies(n2, r2)) // 2
	fmt.Println(maxNumberOfFamilies(n3, r3)) // 5
	fmt.Println(maxNumberOfFamilies(n4, r4)) // 4
	fmt.Println(maxNumberOfFamilies(n5, r5)) // 3
}

// bitwise
// 60 = 1<<2 | 1<<3 | 1<<4 | 1<<5, 2,3,4,5
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	hMap := make(map[int]int)
	for _, r := range reservedSeats {
		hMap[r[0]] |= 1 << r[1]
	}
	var res int
	for _, v := range hMap {
		c := 0
		if ((1<<2 | 1<<3 | 1<<4 | 1<<5) & v) == 0 {
			c++
		}
		if ((1<<6 | 1<<7 | 1<<8 | 1<<9) & v) == 0 {
			c++
		}

		if ((1<<4|1<<5|1<<6|1<<7)&v) == 0 && c == 0 {
			c++
		}
		res += c
	}
	return res + 2*(n-len(hMap))
}

// this solution doesnt work with a cinema with 1 billion rows 10 ^ 9
func maxNumberOfFamiliesOne(n int, reservedSeats [][]int) int {
	a := make([][]int, n)
	var row, fams, seat int
	for i := range a {
		a[i] = make([]int, 10)
		for j := range a[i] {
			a[i][j] = j + 1
		}
	}

	for i := range reservedSeats {
		row = reservedSeats[i][0] - 1
		seat = reservedSeats[i][1] - 1
		a[row][seat] = 0
	}
	for i := range a {
		for j := 9; j > 3; {
			if notReseved(a[i][j-4:j]) && isNotAisle(a[i][j-4:j]) {
				fams++
				j -= 4
			} else {
				j--
			}
		}
	}
	return fams

}

func notReseved(a []int) bool {
	if a[0] != 0 && a[1] != 0 && a[2] != 0 && a[3] != 0 {
		return true
	}
	return false
}

func isNotAisle(a []int) bool {
	s := a[0]
	e := a[3]
	switch {
	case s == 1 && e == 4:
		return false
	case s == 3 && e == 6:
		return false
	case s == 5 && e == 8:
		return false
	case s == 7 && e == 10:
		return false
	}
	return true
}

// attempt with single linear array doesnt work out
func maxNumberOfFamiliesTwo(n int, reservedSeats [][]int) int {
	// 1 make single array mark each reserved seat as 1
	// calculate the space between
	a := make([]int, n*10+n)
	var row, seat, fams, first, last int
	for i := range reservedSeats {
		row = reservedSeats[i][0]
		seat = reservedSeats[i][1]

		if row == 1 {
			a[seat-1] = 1
		} else {
			if seat == 10 {
				seat = row*seat + row - 2
			} else {
				seat = seat + row - 2 + (row*10 - 10)
			}
			a[seat] = 1
		}
	}
	for i := 1; i < n+1; i++ {
		a[i*11-1] = 1
		first = (i - 1) * 11
		last = i*10 + i - 2
		a[first] = 1
		a[last] = 1
	}
	fmt.Println(a)

	for i := len(a) - 1; i > 3; {
		if a[i] == 0 && a[i-1] == 0 && a[i-2] == 0 && a[i-3] == 0 {
			fams++
			a[i], a[i-1], a[i-2], a[i-3] = 1, 1, 1, 1
			i -= 4
		} else {
			i--
		}
	}
	return fams
}
