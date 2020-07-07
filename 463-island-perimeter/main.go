package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(time.Since(t))
}

func islandPerimeter(grid [][]int) int {
	ans := 0
	lenC := len(grid[0]) - 1
	lenR := len(grid) - 1
	for i := range grid {
		for j, tile := range grid[i] {
			if tile == 1 {
				ans += cs(grid, i, j, lenC, lenR)
			}
		}
	}
	return ans
}

func cs(grid [][]int, row, col, lenC, lenR int) int {
	n := 0
	// left
	if col != 0 {
		if grid[row][col-1] == 0 {
			n++
		}
	} else {
		n++
	}
	// right
	if col != lenC {
		if grid[row][col+1] == 0 {
			n++
		}
	} else {
		n++
	}
	// top
	if row != 0 {
		if grid[row-1][col] == 0 {
			n++
		}
	} else {
		n++
	}

	//bottom
	if row != lenR {
		if grid[row+1][col] == 0 {
			n++
		}
	} else {
		n++
	}

	return n
}
