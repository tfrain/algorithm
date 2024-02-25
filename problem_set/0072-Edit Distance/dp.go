// minDistance is a function which calculates the minimum edit distance between two words
func minDistance(word1 string, word2 string) int {
    // Get the lengths of both words
    m, n := len(word1), len(word2)

    // Create a 2D slice for the DP table with rows = word1 length and columns = word2 length
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    // Base case: Initialize the first row (when word2 is empty)
    for i := 1; i <= m; i++ {
        dp[i][0] = i
    }

    // Base case: Initialize the first column (when word1 is empty)
    for i := 1; i <= n; i++ {
        dp[0][i] =  i
    }

    // Helper function to calculate the minimum between two integers
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    // Loop through every character of both words
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                // If characters are equal, cost is 0
                dp[i][j] = dp[i-1][j-1]
            } else {
                // If characters are not equal, cost is minimum edit distance so far + 1
                dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
            }
        }
    }

    // Return the minimum edit distance for the whole words
    return dp[m][n]
}