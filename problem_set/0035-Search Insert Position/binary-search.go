func searchInsert(nums []int, target int) int {
    lo, hi := 0, len(nums)-1  // Initialize the low and high pointers.
    for lo <= hi {  // Run the loop until 'lo' is not greater than 'hi'.
        mid := lo + (hi-lo)/2  // Compute mid index.
        switch {
        case nums[mid] < target:  // If target is greater than mid value, discard left half.
            lo = mid + 1
        case nums[mid] > target:  // If target is smaller than mid value, discard right half.
            hi = mid - 1
        default:
            return mid  // Target found, return mid index.
        }
    }
    // If the target is not present, 'lo' points to the next greater value index where it would be inserted.
    return lo
}

func searchInsert(nums []int, target int) int {
    lo, hi := 0, len(nums)  // Initialize the low and high pointers.
    for lo < hi {  // Run the loop until 'lo' is not equal to 'hi'.
        mid := lo + (hi-lo)/2  // Compute mid index.
        switch {
        case nums[mid] < target:  // If the value is greater than mid value, discard left half.
            lo = mid + 1
        case nums[mid] > target:  // If the value is smaller than mid value, discard right half.
            hi = mid
        default:
            return mid  // Target found, return mid index.
        }
    }
    // If the target is not present, 'lo' points to the next greater value index where it would be inserted.
    return lo
}