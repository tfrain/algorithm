// moveZeroes is a function that takes an array of integers as an argument.
// It moves all the 0's to the end of the array while maintaining the relative order of the non-zero elements.
func moveZeroes(nums []int) {
    // Initialize two pointers i and j to 0
    for i, j := 0, 0; j < len(nums); j++ {
        // If the current element isn't 0
        if nums[j] != 0 {
            // Swap elements at index i and j
            nums[i], nums[j] = nums[j], nums[i]
            // Increment the i pointer
            i++
        }
        // If the current element is 0, j will be incremented in the next iteration
    }
}