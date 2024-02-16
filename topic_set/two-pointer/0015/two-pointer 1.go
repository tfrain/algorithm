func nSumTarget(nums []int, n, start, target int) [][]int { // Define function with input slice 'nums', target numbers 'n', starting index 'start', and target sum 'target'
    sz := len(nums) // Get the length of the nums
    var ret [][]int // Create a variable ret to store the result. The result will be a slice of slices, each inner slice containing n numbers that add up to the target.
    if n < 2 || sz < start { 
        return ret // If n is less than 2 or the starting position exceeds the slice length, return empty slice
    }
    // Two-sum problem
    if n == 2 { // When we're asked for two numbers that add up to 'target'
        lo, hi := start, sz-1 // Initialize two pointers, one at 'start', another at the end of 'nums'
        for lo < hi { // As long as 'lo' pointer doesn't cross 'hi' pointer
            sum := nums[lo] + nums[hi] // Calculate the sum  of the numbers at the 'lo' and 'hi' index
            left, right := nums[lo], nums[hi] // Store the numbers at 'lo' and 'hi' index
            if sum < target { 
                for lo < hi && nums[lo] == left {
                    lo++ // If sum is less than target, move 'lo' pointer to the right
                }
            } else if sum > target {
                for lo < hi && nums[hi] == right {
                    hi-- // If sum is more than target, move 'hi' pointer to the left
                }
            } else {
                ret = append(ret, []int{left, right}) // If sum equals to target, add the pair to the result
                // Skip duplicates
                for lo < hi && nums[lo] == left {
                    lo++
                }
                for lo < hi && nums[hi] == right {
                    hi--
                }
            }
        }
    } else { // For n > 2, recursively reduce the problem to a two-sum problem
        for i := start; i < sz; i++ { 
            sub := nSumTarget(nums, n-1, i+1, target-nums[i]) // Perform recursion to reduce nSum problem by one number
            for _, arr := range sub {
                arr = append(arr, nums[i]) // Append the current number to the list returned by the recursion
                ret = append(ret, arr) // Add the list to the final result
            }
            // Skip duplicates
            for i < sz-1 && nums[i] == nums[i+1] {
                i++
            }
        }
    }
    return ret // Return the final result
}