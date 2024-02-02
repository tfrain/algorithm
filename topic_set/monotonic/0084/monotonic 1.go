func largestRectangleArea(heights []int) int {
    stack := make([]int, 0)  // Initialize the stack.
    maxArea := 0  // Initialize the maximum area as 0.
    for i, h := range heights {
        // Pop the stack once if the current height is smaller than the height of the top of the stack.
        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            //Pop the stack and calculate area.
            j := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1  // Width calculation involving previous stack element.
            }
            area := heights[j] * width  // Calculate area with popped stack height and computed width.
            if area > maxArea {
                maxArea = area  // If area is larger, update maxArea.
            }
        }
        stack = append(stack, i)  // Push current index into the stack.
    }
    // After processing all heights, pop remaining heights from the stack and calculate the area.
    for len(stack) > 0 {
        j := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        width := len(heights)
        if len(stack) > 0 {
            width = len(heights) - stack[len(stack)-1] - 1
        }
        area := heights[j] * width
        if area > maxArea {
            maxArea = area  // If area is larger, update maxArea.
        }
    }
    return maxArea  // Return the maximum possible area.
}