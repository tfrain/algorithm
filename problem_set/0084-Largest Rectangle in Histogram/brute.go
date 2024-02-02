func largestRectangleArea(heights []int) int {
    max := func(a, b int) int {  // Function to get the maximum of two integers.
        if a > b {
            return a
        }
        return b
    }
    var helper func(heights []int, start, end int) int  // Define the helper function.
    helper = func(heights []int, start, end int) int {  // Function to calculate the maximum area.
        if start > end {
            return 0  // If start crosses end, return 0.
        }
        minIndex := start  // Initialize the index of the minimum height.
        for i := start; i <= end; i++ {  
            if heights[i] < heights[minIndex] {  
                minIndex = i  
            }  
        }  
        leftMax := helper(heights, start, minIndex-1)  // Maximum area of left half.
        rightMax := helper(heights, minIndex+1, end)  // Maximum area of right half.
        return max(heights[minIndex]*(end-start+1), max(leftMax, rightMax))  // Maximum area including middle bar.
    }
    return helper(heights, 0, len(heights)-1)  // Return the maximum area.
}