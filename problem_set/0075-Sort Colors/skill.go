func sortColors(nums []int) {
    // Initialize pointers for the three colors
    zero, one, two := -1, 0, len(nums)

    // Loop through the array until the one pointer is less than the two pointer
    for one < two {
        // If the current element is 0, increment the zero pointer and swap the elements at the zero and one pointers
        if nums[one] == 0 {
            zero++
            nums[zero], nums[one] = nums[one], nums[zero]
            one++
        } else if nums[one] == 2 { // If the current element is 2, decrement the two pointer and swap the elements at the two and one pointers
            two--
            nums[two], nums[one] = nums[one], nums[two]
        } else { // If the current element is 1, just increment the one pointer
            one++
        }
    }
}