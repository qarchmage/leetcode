package main

import "fmt"

func main() {

	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}

func groupThePeople(groupSizes []int) [][]int {

	result := [][]int{}

	groupsAndPpls := make(map[int][]int)

	// First we need to build a map where key is capacity of group and value is slice with IDs for this capacity
	for i, v := range groupSizes {
		groupsAndPpls[v] = append(groupsAndPpls[v], i)
	}

	tempGroup := []int{}

	// Then iterate over hashmap for capacity and IDs
	for cap, IDs := range groupsAndPpls {

		// for each ID for this capacity
		for _, v := range IDs {
			// append ID to temporary group
			tempGroup = append(tempGroup, v)
			// if group reachs its limit
			if len(tempGroup) == cap {
				// add this group to the result group and make tempgroup blank again
				result = append(result, tempGroup)
				tempGroup = []int{}
			}
		}
	}

	return result

}
