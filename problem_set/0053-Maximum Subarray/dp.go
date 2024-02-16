func maxSubArray(nums []int) int {
    var maxFunc = func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    pre, max := nums[0], nums[0]
    for _, v := range nums[1:] {
        if pre > 0 {
            pre = pre+v
        } else {
            pre = v
        }
        max = maxFunc(max, pre)
    }
    return max
}

func maxSubArray(nums []int) int {
    dp := make([]int, len(nums))
    max := nums[0]
    dp[0]= nums[0]
    for i := 1; i < len(nums); i++ {
        dp[i] = max(nums[i], nums[i]+dp[i-1])
        if dp[i] > max {
            max = dp[i]
        }
    }
    return max
}