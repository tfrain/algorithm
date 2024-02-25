func trap(height []int) int {
    // Initialize left and right pointers
    left, right := 0, len(height)-1

    // Initialize variables for storing the maximum bar heights on left and right
    // and the resultant water amount
    lMax, rMax, res := 0, 0, 0

    // Helper function to return the maximum of two integers
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }

    // Loop until the pointers meet
    for left <= right {
        // Find maximum bar height from the left and right side
        lMax = max(lMax, height[left])
        rMax = max(rMax, height[right])

        // The trapped water would be decided by the
        // minimum of max height on either left or right
        if lMax < rMax {
            // The difference between maximum bar height and 
            // current bar height will give the trapped water
            // Increment the left pointer
            res += lMax - height[left]
            left++
        } else {
            // Update the resultant water and
            // decrement the right counter
            res += rMax - height[right]
            right--
        }
    }

    // Return the total amount of trapped rain water
    return res
}