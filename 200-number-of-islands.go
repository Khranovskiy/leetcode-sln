package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	result := numIslands(board)
	fmt.Println(result)
}

func numIslands(g [][]byte) int {
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '1' {
				count++
				recursiveDelete(g, i, j)
			}
		}
	}
	return count
}
func recursiveDelete(g [][]byte, i, j int) {
	g[i][j] = '0'
	if j < len(g[i])-1 && g[i][j+1] == '1' {
		recursiveDelete(g, i, j+1)
	}
	if i < len(g)-1 && g[i+1][j] == '1' {
		recursiveDelete(g, i+1, j)
	}
	if j > 0 && g[i][j-1] == '1' {
		recursiveDelete(g, i, j-1)
	}
	if i > 0 && g[i-1][j] == '1' {
		recursiveDelete(g, i-1, j)
	}
}
