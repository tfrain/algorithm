// Function longestValidParentheses calculates the length of the longest valid parentheses substring in string s
func longestValidParentheses(s string) int {
    // Get the length of the string
    n := len(s)
    // Initialize a dynamic programming (dp) array, of length n+1, with all zeros
    dp := make([]int, n+1)
    // Initialize a stack to keep track of the indices of unpaired left parentheses
    stack := make([]int, 0)
    // Initialize a variable to store the maximum length of valid parentheses substring
    res := 0
    // Loop through each character in the string
    for i := range s {
        // Check if the current character is an opening parenthesis
        if s[i] == '(' {
            // If so, set the corresponding dp value to 0, indicating the pairing process has not started yet
            dp[i+1] = 0
            // Add the index of the opening parenthesis to the stack
            stack = append(stack, i)
        } else { 
            // If current character is a closing parenthesis
            // Check whether there's an unpaired opening parenthesis in the stack
            if len(stack) > 0 {
                // Get the index of the last unpaired opening parenthesis
                lastIdx := stack[len(stack)-1]
                // Remove that index from the stack since now it has found its pair
                stack = stack[:len(stack)-1]
                // Check if the length of newly found valid parentheses substring is larger than current max length 'res'
                if res < i-lastIdx+1+dp[lastIdx] {
                    // If so, then update the max length 'res'
                    res = i-lastIdx+1+dp[lastIdx]
                }
                // Record the length of the valid parentheses substring ending at the current position in dp array
                dp[i+1] = i-lastIdx+1+dp[lastIdx]
            } else {
                // If stack is empty, it means we can't form a valid pair with current closing parenthesis so dp value at that position will be 0
                dp[i+1] = 0
            }
        }
    }
    // After finishing scanning the string, 'res' will have the length of the longest valid parentheses substring
    return res
}