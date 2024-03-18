// Top-down approach
func canPartition(nums []int) bool {
    sum, n := 0, len(nums)
    for _, v := range nums {
        sum += v
    }
    if sum%2 != 0 {  // if the sum is odd, we cannot partition it equally
        return false
    } 
    sum = sum / 2  // target sum
    memo := make([][]int, n+1)  // memoization table
    for i := range memo {
        memo[i] = make([]int, sum+1)
    }
    // define dp function for recursion and memoization
    var dp func(amount, start int) bool
    dp = func(amount, start int) bool {
        if amount == 0 {  // if the target sum is reached
            return true
        }
        // if start index is out of bound or the amount is negative, return false
        if start < 0 || start >= n || amount < 0 {
            return false
        }
        // if already computed, return the stored result
        if memo[start][amount] != 0 {
            return memo[start][amount] == 1
        }
        // recursively call dp function: whether to include or exclude the current number
        if dp(amount, start+1) || dp(amount-nums[start], start+1) {
            memo[start][amount] = 1
            return true
        } else {  
            memo[start][amount] = -1
            return false
        }
    }
    return dp(sum, 0)
}


Bottom-up approach
func canPartition(nums []int) bool {
    sum, n := 0, len(nums)
    for _, v := range nums {
        sum += v
    }
    if sum%2 != 0 {
        return false
    }
    sum = sum / 2
    dp := make([][]bool, n+1)  // dp table for storing intermediate results
    for i := range dp {
        dp[i] = make([]bool, sum+1)
        dp[i][0] = true  // it is always possible to sum 0
    }
    for i := 1; i <= n; i++ {
        for j := 1; j <= sum; j++ {
            // if the current number is greater than sum, skip it
            if j-nums[i-1] < 0 {
                dp[i][j] = dp[i-1][j]
            } else {
                // either include the current number nums[i-1] or exclude it
                dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
            }
        }
    }
    return dp[n][sum]
}

// Optimized for space complexity
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum%2 != 0 {
        return false
    }
    sum = sum / 2  // half of total sum
    n := len(nums)
    dp := make([]bool, sum+1)  // reducing the dp table to 1D
    dp[0] = true  // base case: it is always possible to sum 0
    for i := 0; i < n; i++ {
        for j := sum; j >= 0; j-- {
            if j-nums[i] >= 0 {
                // update the dp value
                dp[j] = dp[j] || dp[j-nums[i]]
            }
        }
    }
    return dp[sum]
}