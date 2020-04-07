var maxProfit = function(prices) {
    let total = 0
    let maxIndex = prices.length - 1
    for(let i = 0; i < maxIndex; i++) {
      if(prices[i+1] > prices[i]){
        total += prices[i+1] - prices[i]
      }
    }
    return total;
}
