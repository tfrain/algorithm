// lengthOfLIS calculates the length of the longest increasing subsequence in an array.
func lengthOfLIS(nums []int) int {
    // Variable to store the result - the length of the longest subsequence.
    var res int
    // dp array where dp[i] will store the length of the longest increasing subsequence ending with nums[i].
    dp := make([]int, len(nums))
    // Initialize each subsequence length to 1, as each element is a subsequence of length 1.
    for i := range dp {
        dp[i] = 1
    }

    // Helper function to find the maximum of two integers.
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }

    // Loop through the array to build up the dp array.
    for i, v := range nums {
        // Check all previous elements to find increasing subsequences ending with v.
        for j := 0; j < i; j++ {
            // If v is greater than a previous number, update the dp value for v.
            if v > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        // Update the result if the current subsequence is the longest so far.
        if dp[i] > res {
            res = dp[i]
        }
    }
    // Return the length of the longest increasing subsequence.
    return res
}