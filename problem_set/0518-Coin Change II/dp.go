func change(amount int, coins []int) int {
    n := len(coins)
    dp := make([][]int, n+1) // Create a 2D dp array of size (n+1)x(amount+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, amount+1) // Initialize the DP array
        dp[i][0] = 1 // Set the base case
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= amount; j++ {
            if j - coins[i-1] >= 0 {
                dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]] // Calculate the ways
            } else {
                dp[i][j] = dp[i-1][j] // Copy the previous row value
            }
        }
    }
    return dp[n][amount] // Return the number of ways to make the amount
}

func change(amount int, coins []int) int {
    n := len(coins)
    dp := make([]int, amount+1) // Create a 1D dp array of size (amount+1)
    dp[0] = 1 // base case
    for i := 0; i < n; i++ {
        for j := 1; j <= amount; j++ {
            if j-coins[i] >= 0 {
                dp[j] = dp[j] + dp[j-coins[i]] // Calculate the ways
            }
        }
    }
    return dp[amount] // Return the number of ways to make the amount
}