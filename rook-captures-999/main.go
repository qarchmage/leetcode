package main

import "fmt"

func main() {
	a := []string{"........", "...p....", "...R...p", "........", "........", "...p....", "........", "........"}
	board := make([][]byte, 8, 8)
	for i := range a {
		board[i] = []byte(a[i])
	}
	//for i := range board {
	//fmt.Println(len(board[i]))
	//}
	//fmt.Println(board)
	fmt.Println(numRookCaptures(board))

}
func numRookCaptures(board [][]byte) int {
	var ans, x, y int
	// find rook
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'R' {
				x = i
				y = j
			}
		}
	}

	// search top , right ,bottom, left until B , end,p
	// top
	for i := x - 1; i > -1; i-- {
		if board[i][y] == 'p' {
			ans++
			break
		}
		if board[i][y] == 'B' {
			break
		}
	}
	// right
	for i := y; i < 8; i++ {
		if board[x][i] == 'p' {
			ans++
			break
		}
		if board[x][i] == 'B' {
			break
		}
	}
	//bottom
	for i := x; i < 8; i++ {
		if board[i][y] == 'p' {
			ans++
			break
		}
		if board[i][y] == 'B' {
			break
		}
	}
	// left
	for i := y; i > -1; i-- {
		if board[x][i] == 'p' {
			ans++
			break
		}
		if board[x][i] == 'B' {
			break
		}
	}

	return ans

}
