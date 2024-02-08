func findDuplicate(nums []int) int {
  
    // We start both pointers from the first number in the array.
    slow, fast := nums[0], nums[nums[0]]

    // The `slow` pointer makes one step, the `fast` pointer makes two steps.
    // We continue this process until both pointers meet inside the cycle.
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }

    // We restart the `fast` pointer from the first number in the array.
    fast = 0

    // Both pointers make one step at a time.
    // We continue this process until both pointers meet at the start of the cycle, which is the duplicate number.
    for fast != slow {
        slow = nums[slow]
        fast = nums[fast]
    }

    // The duplicate number is returned.
    return slow
}