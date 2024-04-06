func canJump(nums []int) bool {
    fastest := 0
    // Loop over all but last element
    for i := 0; i < len(nums)-1; i++ {
        // Keep track of maximum jump
        if fastest < nums[i]+i {
            fastest = nums[i] + i
        }
        // If max jump is still not sufficient to cross current index, return false
        if fastest <= i {
            return false
        }
    }
    // We are always able to reach the last index
    return true
}