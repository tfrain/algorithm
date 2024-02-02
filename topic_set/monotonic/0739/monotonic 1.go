func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    res, stack := make([]int, n), []int{}
    for i := n-1; i >= 0; i-- {
        // Push temperatures onto the stack that are higher
        // Maintain decreasing order of temperatures in the stack
        for len(stack) != 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
            stack = stack[:len(stack)-1]
        }
        // If stack is empty, res[i] stays 0, else we calculate the warmer day distance
        if len(stack) != 0 {
            res[i] = stack[len(stack)-1] - i 
        }
        // Push current day into the stack
        stack = append(stack, i) 
    }
    return res // return the result
}