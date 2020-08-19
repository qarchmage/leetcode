package main

import "fmt"

func main() {

	t := [][]int{{2, 6, 4, 1}, {1, 2, 34, 3, 4, 5, 7, 23, 12}}

	for _, v := range t {
		fmt.Println(threeConsecutiveOdds(v))
	}

}

func threeConsecutiveOdds(arr []int) bool {
    for i := 2; i < len(arr); i++ {
        if arr[i] % 2 == 1 && arr[i-1] % 2 == 1 && arr[i-2] % 2 == 1 {
            return true
        }
    }
    return false
}

