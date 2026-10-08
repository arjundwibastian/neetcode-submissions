class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices) {
        let profit = 0;
        let currentProfit;
        for (let i = 0; i < prices.length; i++) {
            for (let j = 0; j < prices.length; j++) {
                if (j > i) {
                currentProfit = prices[j] - prices[i];
                if (currentProfit > profit) profit = currentProfit;
                }
             }
             }
        return profit > 0 ? profit : 0;
    }
}
