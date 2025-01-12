package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int = 2
	fmt.Println(n / 3)

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	var (
		rows    [9][9]bool
		columns [9][9]bool
		quads   = make(map[string][9]bool, 9)
	)
	n := len(board)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == '.' {
				continue
			}
			strNum := string(board[row][col])
			num, _ := strconv.Atoi(strNum)
			pos := num - 1
			key := fmt.Sprintf("%d:%d", row/3, col/3)

			if rows[row][pos] || columns[col][pos] || quads[key][pos] {
				return false
			}
			rows[row][pos] = true
			columns[col][pos] = true

			temp := quads[key]
			temp[pos] = true
			quads[key] = temp
		}
	}
	return true
}
