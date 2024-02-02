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
        if v-min > 0 {
            profit += v - min
            min = v
        }
    }
    // Return the maximum profit made.
    return profit
}