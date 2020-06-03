package main

import (
	"fmt"
	"sort"
)

func main() {
	a := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	fmt.Println(twoCitySchedCost(a))
}

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	result := 0
	for index, cost := range costs {
		if index < len(costs)/2 {
			result += cost[0]
		} else {
			result += cost[1]
		}
	}
	return result
}
