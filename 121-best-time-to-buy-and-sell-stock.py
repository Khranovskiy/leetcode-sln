class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        maxProfit = 0
        bestLowPrice = prices[0]

        for price in prices[1:]:
            maxProfit = max(maxProfit, price - bestLowPrice)
            bestLowPrice = min(bestLowPrice, price)
        return maxProfit

  
