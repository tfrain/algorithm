func firstMissingPositive(nums []int) int {
    // we iterate through the array
    for i := 0; i < len(nums); i++ {
        // we check if the current number is in the range of 1 to the length of the array
        // we also check if the number at the target position is not already the current number
        // to prevent infinite swapping
        if nums[i] >= 1 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
            // we put the current number to its correct position, which is nums[i]-1 index in the array
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
            // we decrement i to stay at the current position
            // because the swapped number could also be a valid positive number that has not been processed
            i--
        }
    }
    // we look for the first missing positive number
    for i := range nums {
        if nums[i] != i+1 {
            return i+1
        }
    }
    // if all numbers are present, we return the next positive number after the length of the array
    return len(nums)+1
}

func firstMissingPositive(nums []int) int {
    // we iterate through the array
    for i := range nums {
        // we keep swapping the current number to its correct position, until the current number is
        // out of range or is already at its correct position
        for j := nums[i] - 1; j >= 0 && j < len(nums); j = nums[i] - 1 {
            if i == j || nums[i] == nums[j] {
                break
            }
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    // we check from the beginning of the array for the first number that is not in its correct position
    var i int
    for i < len(nums) && nums[i] == i+1 {
        i++
    }
    // we return the first missing positive number
    return i + 1
}