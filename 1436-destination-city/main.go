package main

import "fmt"

func main() {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	paths2 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}

	paths3 := [][]string{{"qMTSlfgZlC", "ePvzZaqLXj"}, {"xKhZXfuBeC", "TtnllZpKKg"}, {"ePvzZaqLXj", "sxrvXFcqgG"}, {"sxrvXFcqgG", "xKhZXfuBeC"}, {"TtnllZpKKg", "OAxMijOZgW"}}
	fmt.Println(destCity(paths))
	fmt.Println(destCity(paths2))
	fmt.Println(destCity(paths3))
}

func destCity(paths [][]string) string {
	r := paths[0][1]
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths); j++ {
			if r == paths[j][0] {
				r = paths[j][1]
			}
		}
	}

	return r
}
