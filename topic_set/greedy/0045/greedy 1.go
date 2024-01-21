func jump(nums []int) int {
    // Initialize required variables and functions
    jumps, end, fastest := 0, 0, 0
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    for i := 0; i < len(nums)-1; i++ {
        // Find the fastest reachable position
        fastest = max(fastest, nums[i]+i)
        if i == end {
            // Update the end position and increment the number of jumps
            jumps++
            end = fastest
        }
    }
    // Return the minimum number of jumps to reach the end
    return jumps
}