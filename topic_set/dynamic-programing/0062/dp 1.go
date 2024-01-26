func uniquePaths(m int, n int) int {
    // Initialize DP table
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][0] = 1
    }
    // First row is 1 as there's only one way to move right
    for i := 1; i < n; i++ {
        dp[0][i] = 1
    }
    // For the rest cells, sum up paths from top and left
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    return dp[m-1][n-1]
}

func uniquePaths(m int, n int) int {
    // Initialize DP memo
    memo := make([][]int, m)
    for i := range memo {
        memo[i] =  make([]int, n)
    }
    // Recursive function to compute paths
    var dp func(x, y int) int
    dp = func(x, y int) int {
        // Base case: start position is 1 path itself
        if x == 0 && y == 0 {
            return 1
        }
        // Going out of grid isn't a valid path
        if x < 0 || y < 0 {
            return 0
        }
        // If we've computed paths for this cell already, use it
        if memo[x][y] > 0 {
            return memo[x][y]
        }
        // Recursive call to tracks from top and left
        memo[x][y] = dp(x-1, y) + dp(x, y-1)
        return memo[x][y]
    }
    // We start from the finish(bottom-right) and make our way to start
    return dp(m-1, n-1)
}