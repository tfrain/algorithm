// Function to find the minimum element in a rotated sorted array
func findMin(nums []int) int {
    // Initial pointers setup: l to start of the array, 
    // r to end of the array
    l, r := 0, len(nums)-1
    
    // Iterate while left pointer is less than right pointer
    for l < r {
        // Calculate the middle index using integer division
        m := (l + r) / 2
        
        // If middle element > the next element ("m+1"), 
        // return the next element
        if nums[m] > nums[m+1] {
            return nums[m+1]
        }
        // If middle element > last element, 
        // move the left pointer to m + 1, since min lies to the right of m
        if nums[m] > nums[r] {
            l = m + 1
        } else {
            // Else, move the right pointer to m, 
            // since min lies to its left
            r = m
        }
    }
    // Return value at l, since it's pointing to the minimum element
    return nums[l]
}

// Function to find the minimum element in a rotated sorted array
func findMin(nums []int) int {
    // Initial pointers setup: l to start of the array, 
    // r to end of the array
    l, r := 0, len(nums)-1
    
    // Iterate while left pointer is less than right pointer
    for l < r {
        // Calculate the middle index
        m := l + (r-l)/2
        
        // If middle element is greater or equal to first element,
        // move the left pointer to m + 1
        // This means rotation might be to the right of m
        if nums[m] >= nums[0] {
            l = m + 1
        } else {
            // If middle < middle - 1, return the middle element
            if nums[m] < nums[m-1] {
                return nums[m]
            } else {
                // If else, move the right pointer
                r = m - 1
            }
        }
    }
    // If left pointer is at the end and element at l > element at 0,
    // array rotation was at beginning
    if l == len(nums)-1 && nums[l] > nums[0] {
        l = 0
    }
    return nums[l]
}