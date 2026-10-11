func maxProfit(prices []int) int {
    lowestPrice := prices[0]
    var maxProfit int
    for _, price := range prices {
        if price < lowestPrice {
            lowestPrice = price
            continue
        } else if price > lowestPrice {
            currentProfit := price - lowestPrice
            if maxProfit < currentProfit {
                maxProfit = currentProfit
            } 
        }
    }
    return maxProfit
}
