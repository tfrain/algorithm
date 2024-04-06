func canJump(nums []int) bool {
    // Get length of nums
    n := len(nums)
    // Initialize Memoization Table
    memo := make([]int, n)
    var dp func(start int) bool
    dp = func(start int) bool {
        // If start is greater than or equal to last index, return true
        if start >= n-1 {
            return true
        }
        // If we have already calculated the possible jump at this start point, reuse
        if memo[start] != 0 {
            return memo[start] == 1
        }
        // The max jump we can make at this point
        maxJump := nums[start]
        for i := 1; i <= maxJump; i++ {
            if dp(start + i) {
                memo[start] = 1
                return true
            }
        }
        memo[start] = -1
        return false
    }
    // Perform recursive dp starting from first index
    return dp(0)
}

func canJump(nums []int) bool {
    n := len(nums)
    // Initialize dp table
    dp := make([]bool, n)
    dp[n-1] = true
    // Start from end
    for i := n - 2; i >= 0; i-- {
        // Try for each possible jump
        for j := i; j <= i+nums[i] && j < n; j++ {
            dp[i] = dp[i] || dp[j]
        }
    }
    return dp[0]
}
