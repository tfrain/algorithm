func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    // s1[0..i-1] and s2[0..j-1]'s LCS length is dp[i][j]
    dp := make([][]int, m+1)
    // target: s1[0..m-1] and s2[0..n-1]'s LCS length, which is dp[m][n]
    // base case: dp[0][..] = dp[..][0] = 0
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            // now i and j start from 1, so subtract 1
            if text1[i-1] == text2[j-1] {
                // s1[i-1] and s2[j-1] must be in LCS
                dp[i][j] = dp[i-1][j-1]+1
            } else {
                // s1[i-1] and s2[j-1] at least one is not in LCS
                dp[i][j] = max(dp[i][j-1], dp[i-1][j])
            }
        }
    }
    return dp[m][n]
}