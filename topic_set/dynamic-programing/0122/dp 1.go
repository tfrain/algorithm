func maxProfit(prices []int) int {
    // Initializing profit variable to store total profit.
    var profit int
    // Variable 'min' is initialized to MaxInt32 to capture minimum price of stock.
    min := math.MaxInt32
    for _, v := range prices {
        // Update 'min' if we find a price less than 'min'.
        if v < min {
            min = v
        }
        // Check if selling the stock today would yield profit. If yes, 
        // update profit and set 'min' to today's price for next transaction.
        if v-min > 0 {afunc maxProfit(prices []int) int {
    n := len(prices)
    // Array 'dp' of size 'n' is initialized to hold profit for each day.
    dp := make([][]int, n)
    max := func(a, b int) int {
        // Using 'max' function to compare and return the maximum value.
        if a > b {
            return a
        }
        return b
    }
    for i := 0; i < n; i++ {
        // Two-dimensional 'dp' array,
        // where 'dp[i][0]' represents maximum profit on day 'i' when we do not hold the stock,
        // 'dp[i][1]' represents maximum profit on day 'i' when we hold the stock.
        dp[i] = make([]int, 2)
        if i == 0 {
            // Base case: profit is zero on the first day if we do not hold the stock, 
            // and negative of the price of the stock if we buy it on the first day.
            dp[i][0] = 0
            dp[i][1] = -prices[i]
            continue
        }
        // Update 'dp[i][0]' and 'dp[i][1]' based on profits on previous day.
        dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
        dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
    }
    // Return the maximum profit on the last day when we do not hold the stock.
    return dp[n-1][0]
}
            profit += v - min
            min = v
        }
    }
    // Return the maximum profit made.
    return profit
}