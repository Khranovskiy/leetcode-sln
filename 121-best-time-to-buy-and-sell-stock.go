func maxProfit(prices []int) int {
    maxProfit := 0
    bestLowPrice := prices[0]

    for _, price := range prices[1:] {
        maxProfit = max(maxProfit, price - bestLowPrice)
        bestLowPrice = min(bestLowPrice, price)
    }

    return maxProfit
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
