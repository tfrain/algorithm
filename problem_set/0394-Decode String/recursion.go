func decodeString(s string) string {
    n := len(s)
    var helper func(start int) (string, int)
    helper = func(start int) (string, int) {
        // Base case: if start >= n
        if start >= n {
            return "", 0
        }
        // Initialize variables for current substring and current index
        var cur string
        i, num := start, 0
        // Start of main processing loop
        for i < n {
            // If the current character is a number
            if v, err := strconv.Atoi(string(s[i])); err == nil {
                num = num*10+v
            } else if s[i] == '[' { // If the current character is an opening bracket
                // Apply recursion to handle next level
                sub, end := helper(i+1)
                // Append the decoded substring 'num' times to 'cur'
                cur += strings.Repeat(sub, num)
                i = end
                num = 0 // Reset 'num' for the next number-bracket set in string.
            } else if s[i] == ']' { // If the current character is a closing bracket
                return cur, i
            } else { // If the current character is a letter
                cur += string(s[i])
            }
            i++
        }
        // Return the final decoded string and the final position
        return cur, i
    }
    // Run the helper function starting from position 0
    res, _ := helper(0)
    return res
}