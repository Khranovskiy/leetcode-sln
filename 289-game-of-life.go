package main

import (
	"fmt"
)

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}

const (
	dead   = 0
	live   = 1
	toLive = 2
	toDead = 3
)

func gameOfLife(b [][]int) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			c := getliveNeighbors(i, j, b)

			if b[i][j] == dead {
				if c == 3 {
					b[i][j] = toLive
				}
				continue
			}

			if c < 2 || c > 3 {
				b[i][j] = toDead
				continue
			}

			b[i][j] = live
		}
	}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == toDead {
				b[i][j] = dead
			}

			if b[i][j] == toLive {
				b[i][j] = live
			}
		}
	}
}

func getliveNeighbors(i, j int, board [][]int) int {
	up, down, in := 0, 0, 0

	if i > 0 {
		up = getValue(board[i-1][j])

		if j > 0 {
			up += getValue(board[i-1][j-1])
		}

		if j < len(board[i])-1 {
			up += getValue(board[i-1][j+1])
		}
	}

	if i < len(board)-1 {
		down = getValue(board[i+1][j])

		if j > 0 {
			up += getValue(board[i+1][j-1])
		}

		if j < len(board[i])-1 {
			up += getValue(board[i+1][j+1])
		}
	}

	if j > 0 {
		in += getValue(board[i][j-1])
	}

	if j < len(board[i])-1 {
		in += getValue(board[i][j+1])
	}

	return up + down + in
}

func getValue(val int) int {
	if val == 2 {
		return 0
	}

	if val == 3 {
		return 1
	}

	return val
}
