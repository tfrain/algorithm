func jump(nums []int) int {
    // Initialize required variables and functions
    n := len(nums)
    memo := make([]int, n)
    for i := range memo {
      memo[i] = n-1
    }
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    var dp func(start int) int
    // Define the dp function
    dp = func(start int) int {
        // If it is the end or beyond, then no jump is needed
        if start >= n-1 {
            return 0
        }
        // Check if we already have a solution
        if memo[start] != n-1 {
            return memo[start]
        }
        // Recursive step, try all possible steps that we can jump to
        for i := 1; i <= nums[start]; i++ {
            tmp := dp(start+i)
            // Minimize the number of jumps and save it in memo
            memo[start] = min(memo[start], tmp+1)
        }
        return memo[start]
    }
    // Kick-off the function with the first position
    return dp(0)
}

func jump(nums []int) int {
    // Initialize required variables and functions
    n := len(nums)
    dp := make([]int, n)
    for i := range dp {
  dp[i] = n-1
 }
 dp[n-1] = 0
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }
    for i := n-2; i >= 0; i-- {
        // Try all reachable positions
        for j := i+1; j <= nums[i]+i && j < n; j++ {
            // Minimize the number of jumps
            dp[i] = min(dp[i], dp[j]+1)
        }
    }
    // Return the minimum number of jumps to reach the first position
    return dp[0]
}