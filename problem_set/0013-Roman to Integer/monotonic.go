
// Function to convert Roman numerals to integer
func romanToInt(s string) int {
    // Create a map to store the integer values of Roman numerals
    m := make(map[byte]int)
    m['I'] = 1
    m['V'] = 5
    m['X'] = 10
    m['L'] = 50
    m['C'] = 100
    m['D'] = 500
    m['M'] = 1000

    // Create a stack to store the indices of the string
    stack := make([]int, 0)
    res := 0

    // Iterate over the string
    for i := range s {
        // If the stack is not empty and the current numeral is smaller or equal to the last numeral in the stack
        if len(stack) > 0 && m[s[stack[len(stack)-1]]] >= m[s[i]] {
            // Compute the integer value of the Roman numerals in the stack and add it to the result
            res += compute(m, s[stack[0]:stack[len(stack)-1]+1])
            // Clear the stack
            stack = stack[:0]
        }
        // Add the current index to the stack
        stack = append(stack, i)
    }

    // If the stack is not empty after iterating over the string
    if len(stack) > 0 {
        // Compute the integer value of the remaining Roman numerals in the stack and add it to the result
        res += compute(m, s[stack[0]:stack[len(stack)-1]+1])
    }

    // Return the result
    return res
}

// Function to compute the integer value of a string of Roman numerals
func compute(m map[byte]int, s string) int {
    // Initialize the result with the value of the last numeral
    res := m[s[len(s)-1]]
    // Iterate over the string in reverse order
    for i := len(s) - 2; i >= 0; i-- {
        // Subtract the value of the current numeral from the result
        res -= m[s[i]]
    }
    // Return the result
    return res
}