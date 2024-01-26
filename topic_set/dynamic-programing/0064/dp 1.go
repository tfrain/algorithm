// minPathSum calculates the minimum path sum from the top-left corner to the bottom-right corner of a grid.
func minPathSum(grid [][]int) int {
    // Get the dimensions of the grid.
    m, n := len(grid), len(grid[0])

    // Initialize the DP table with the same dimensions as the grid.
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    // Helper function to find the minimum of two integers.
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    // Initialize the starting point of the DP table with the grid's starting point value.
    dp[0][0] = grid[0][0]

    // Initialize the first column of the DP table.
    for i := 1; i < m; i++ {
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }

    // Initialize the first row of the DP table.
    for i := 1; i < n; i++ {
        dp[0][i] = dp[0][i-1] + grid[0][i]
    }

    // Populate the rest of the DP table.
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            // The value in dp[i][j] is the sum of the current grid value and the minimum of the
            // two possible previous steps (from the top or from the left).
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }

    // The last element in the DP table contains the minimum path sum.
    return dp[m-1][n-1]
}

func minPathSum(grid [][]int) int {
    // Get the dimensions of the grid.
    m, n := len(grid), len(grid[0])

    // Initialize a memoization table with the same dimensions as the grid.
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        // Set initial values to -1 to indicate that no calculation has been done for this cell.
        for j := 0; j < n; j++ {
            memo[i][j] = -1
        }
    }

    // Helper function to find the minimum of two integers.
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    // dp is a recursive function that computes the minimum path sum to reach cell (i, j) from (0, 0).
    var dp func(grid [][]int, i, j int) int
    dp = func(grid [][]int, i, j int) int {
        // Base case: return the starting point of the grid.
        if i == 0 && j == 0 {
            return grid[0][0]
        }
        // If out of bounds, return the maximum integer value as this path is not viable.
        if i < 0 || j < 0 {
            return math.MaxInt32
        }
        // If the value is already computed, return it from the memo table.
        if memo[i][j] != -1 {
            return memo[i][j]
        }
        // Recursively compute the minimum path sum using the memo table to store results.
        memo[i][j] = min(dp(grid, i-1, j), dp(grid, i, j-1)) + grid[i][j]
        return memo[i][j]
    }

    // Compute the minimum path sum for the bottom-right corner.
    return dp(grid, m-1, n-1)
}