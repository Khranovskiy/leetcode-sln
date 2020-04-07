package main

func maxProfit(prices []int) int {
	total := 0
	maxIndex := len(prices) - 1
	for i := 0; i < maxIndex; i++ {
		if prices[i+1] > prices[i] {
			total += prices[i+1] - prices[i]
		}
	}
	return total
}
