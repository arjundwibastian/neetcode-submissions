func maxProfit(prices []int) int {
	lowestPrice := 0
    var maxProfit int
    for i := 0; i < len(prices); i++{
        if i == 0 {
            lowestPrice = prices[i]
            continue
        }
        if lowestPrice > prices[i] {
            lowestPrice = prices[i]
            continue
        }
        currentProfit := prices[i] - lowestPrice
        if maxProfit < currentProfit {
            maxProfit = currentProfit
        }
    }
    return maxProfit
}
