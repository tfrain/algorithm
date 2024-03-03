func searchRange(nums []int, target int) []int {
    // Define a method to find first position
    firstPosition := func() int {
        lo, hi := 0, len(nums)
        for lo < hi { // While start pointer is less than end pointer
            m := lo + (hi-lo)/2 // Get mid point
            // Depending on comparison with target, adjust lo and hi pointers
            if nums[m] < target { 
                lo = m + 1
            } else if nums[m] > target {
                hi = m
            } else {
                hi = m
            }
        }
        if lo >= len(nums) || nums[lo] != target {
            return -1 // If no match found, return -1
        }
        return lo
    }
    // Define a method to find last position
    lastPosition := func() int {
        lo, hi := 0, len(nums)
        for lo < hi { // While start pointer is less than end pointer
            m := lo + (hi-lo)/2 // Get mid point
            // Depending on comparison with target, adjust lo and hi pointers
            if nums[m] < target {
                lo = m + 1
            } else if nums[m] > target {
                hi = m
            } else {
                lo = m + 1
            }
        }
        if hi-1 < 0 || nums[hi-1] != target {
            return -1 // If no match found, return -1
        }
        return hi - 1
    }
    return []int{firstPosition(), lastPosition()} // Return array of first and last positions
}

func searchRange(nums []int, target int) []int {
    // Check for edge case of empty array
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    // Define a method for a left-to-right binary search
    leftRound := func(nums []int, target int) int {
        left, right := 0, len(nums)-1
        for left <= right {
            mid := left + (right-left)/2 // Get mid point
            // Depending on comparison with target, adjust left and right pointers
            if nums[mid] < target {
                left = mid + 1
            } else if nums[mid] >= target {
                right = mid - 1
            }
        }
        // If there is no element matching target, return -1
        if left >= len(nums) || nums[left] != target {
            return -1
        }
        return left
    }
    // Define a method for a right-to-left binary search
    rightRound := func(nums []int, target int) int {
        left, right := 0, len(nums)-1
        for left <= right { // While start pointer is less than end pointer
            mid := left + (right-left)/2 // Get mid point
            // Depending on comparison with target, adjust left and right pointers
            if nums[mid] <= target {
                left = mid + 1
            } else if nums[mid] > target {
                right = mid - 1
            }
        }
        // If there is no element matching target, return -1
        if right < 0 || nums[right] != target {
            return -1
        }
        return right
    }
    return []int{leftRound(nums, target), rightRound(nums, target)} // Return array of first and last positions
}