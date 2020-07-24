package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})
	indexes := make([]int, len(people))
	for i := range indexes {
		indexes[i] = i
	}
	queued := make([][]int, len(people))
	for _, v := range people {
		queued[indexes[v[1]]] = v
		indexes = append(indexes[:v[1]], indexes[v[1]+1:]...)
	}
	return queued
}
