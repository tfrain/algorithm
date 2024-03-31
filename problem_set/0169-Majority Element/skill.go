// majorityElement function takes an array of integers as input and returns the majority element.
// The majority element is the element that appears more than ⌊n/2⌋ times in the array, where n is the length of the array.
func majorityElement(nums []int) int {
    // Initialize target as the first element of the array and count as 1
    target, cnt := nums[0], 1

    // Iterate over the array starting from the second element
    for i := 1; i < len(nums); i++ {
        // If the current element is equal to the target, increment the count
        if nums[i] == target {
            cnt++
        } else {
            // If the current element is not equal to the target, decrement the count
            cnt--
        }
        // If the count becomes zero, update the target to the current element and reset the count to 1
        if cnt == 0 {
            target = nums[i]
            cnt=1
        }
    }
    // Return the majority element
    return target
}