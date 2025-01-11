// 064. minimum_path_sum
// https://leetcode.com/problems/minimum-path-sum/
package main

import "fmt"

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

func minPathSum(g [][]int) int {
	var i, j int
	for i = 0; i < len(g); i++ {
		for j = 0; j < len(g[i]); j++ {
			//  x - - -
			//  - - - -
			//  - - - -
			if i == 0 && j == 0 {
				continue
			}
			//  p c x x
			//  - - - -
			//  - - - -
			if i == 0 {
				p := g[i][j-1]
				g[i][j] = p + g[i][j]
				continue
			}
			//  p - - -
			//  c - - -
			//  x - - -
			if j == 0 {
				p := g[i-1][j]
				g[i][j] = p + g[i][j]
				continue
			}
			//  -  p1 - -
			//  p2 c  x x
			//  -  x  x x
			p2, p1 := g[i][j-1], g[i-1][j]
			minPrices := p2
			if p1 < minPrices {
				minPrices = p1
			}
			g[i][j] = minPrices + g[i][j]
		}
	}
	// i == len(g), j = len(g[i])
	return g[i-1][j-1]
}
