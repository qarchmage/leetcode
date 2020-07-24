package main

import "fmt"

func main() {

	tc := [][]int{{9, 3}, {15, 4}}
	for _, v := range tc {
		fmt.Println(numWaterBottles(v[0], v[1]))
	}
}

func numWaterBottles(numBottles int, numExchange int) int {
	var cnt, empty int
	drink := numBottles
	total := numBottles

	if numBottles < numExchange {
		cnt = numBottles
	}
	for total >= numExchange {
		cnt += drink
		total = drink + empty
		drink = total / numExchange
		empty = total % numExchange
	}
	return cnt

}
