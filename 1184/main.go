package main

import "fmt"

func main() {
	fmt.Println(distanceBetweenBusStops([]int{8, 11, 6, 7, 10, 11, 2}, 0, 3))       // 25
	fmt.Println(distanceBetweenBusStops([]int{7, 10, 1, 12, 11, 14, 5, 0}, 7, 2))   // 17
	fmt.Println(distanceBetweenBusStops([]int{3, 6, 7, 2, 9, 10, 7, 16, 11}, 6, 2)) // 28
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if len(distance) == 0 || start == destination {
		return 0
	}

	if start > destination {
		start, destination = destination, start
	}

	clockwise := sum(distance[start:destination])
	counterClockwise := sum(distance) - clockwise
	if clockwise < counterClockwise {
		return clockwise
	}
	return counterClockwise
}

func sum(ar []int) int {
	var sum int
	for _, v := range ar {
		sum += v
	}
	return sum
}
