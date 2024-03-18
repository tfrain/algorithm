func search(nums []int, target int) int {
    // Set starting search boundaries.
    l, r := 0, len(nums)-1

    // Continue searching while boundaries are valid.
    for l <= r {
        // Calculate middle point of search boundary.
        m := l+(r-l) / 2

        // Case where middle value is less than the target value.
        if nums[m] < target  {
            // If middle value is less than the first value and target is greater than or equal to the first value.
            // This means target must lie on the left half. So we move right pointer (r) to left half by assigning it to m-1
            if nums[m] < nums[0] && target >= nums[0] {
                r = m-1
            } else {
                // Moves the left pointer (l) to right half as target cannot be in left half.
                l = m+1
            }
        } else if nums[m] > target {
            // Case where middle value is greater than the target value.
            // If target value is less than the first value and middle value is greater than or equal to the first value.
            // This would mean target lies in the right half. So we assign left pointer to m+1
            if target < nums[0] && nums[m] >= nums[0] {
                l = m+1
            } else {
                // Moves the right pointer (r) to left half as target cannot be in right half.
                r = m-1
            }
        } else {
            // Target is found at position m.
            return m
        } 
    }
    // Target not found so return -1.
    return -1
}