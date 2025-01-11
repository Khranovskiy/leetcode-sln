package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(p []int) int {
	// [7,1,5,3,6,4]
	// You can only hold at most one share of the stock at any time
	// However, you can buy it then immediately sell it on the same day.
	// like skip a day

	// decide to buy and/or sell the stock

	// Find and return the maximum profit you can achieve

	//[7,1,5,3,6,4]      B buy
	// - B S B S -  x?   - skip (buy+sell)
	//     4   3    7    S sell

	res, d := 0, 0
	l := len(p)
	for i := 0; i < l-1; i++ {
		d = p[i+1] - p[i]

		if d > 0 {
			res = res + d
		}
	}
	return res
}

// 4min 30 sec
