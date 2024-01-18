// findKthLargest finds the k-th largest element in an unsorted array using QuickSelect.
func findKthLargest(nums []int, k int) int {
    n := len(nums)
    // Position of the k-th largest element in sorted order.
    pos := n - k
    // Initialize left and right pointers.
    l, r := 0, len(nums)-1
    // QuickSelect algorithm.
    for l <= r {
        // Partition the array and get the index of the pivot.
        idx := partition(nums, l, r)
        // If the pivot is in the k-th position, return its value.
        if idx == pos {
            return nums[idx]
        // If the pivot is to the right of the k-th position, move left.
        } else if idx > pos {
            r = idx - 1
        // If the pivot is to the left of the k-th position, move right.
        } else {
            l = idx + 1
        }
    }
    // If we don't find the k-th largest, return -1.
    return -1
}

// partition reorders the elements in nums[l...r] so that all elements less than the pivot
// come before the pivot, while all elements greater than the pivot come after it.
func partition(nums []int, l, r int) int {
    // Choose the leftmost element as the pivot.
    pivot := nums[l]
    // Initialize pointers for the subarray to partition.
    left, right := l+1, r
    // Partition the array based on pivot value.
    for left <= right {
        // Move left pointer to the right as long as nums[left] is less than or equal to the pivot.
        for left <= right && nums[left] <= pivot {
            left++
        }
        // Move right pointer to the left as long as nums[right] is greater than the pivot.
        for left <= right && nums[right] > pivot {
            right--
        }
        // Swap elements to ensure all nums less than pivot are to its left, and all nums greater
        // are to its right.
        if left <= right {
            nums[left], nums[right] = nums[right], nums[left]
            left++
            right--
        }
    }
    // Swap the pivot with nums[right] so that pivot is in its final sorted position.
    nums[right], nums[l] = nums[l], nums[right]
    // Return the index of the pivot.
    return right
}