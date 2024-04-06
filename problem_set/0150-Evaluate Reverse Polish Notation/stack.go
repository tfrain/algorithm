// evalRPN function evaluates the value of an arithmetic expression in Reverse Polish Notation.
func evalRPN(tokens []string) int {
    // Initialize an empty stack to store the numbers.
    stk := make([]int, 0)

    // Iterate over each token in the input array.
    for _, token := range tokens {
        // If the token is an operator, perform the operation.
        if strings.Contains("+-*/", token) {
            // Pop the last two numbers from the stack.
            a, b := stk[len(stk)-1], stk[len(stk)-2]
            // Remove the last two numbers from the stack.
            stk = stk[:len(stk)-2]

            // Perform the operation based on the operator and push the result back to the stack.
            switch token {
            case "+":
                stk = append(stk, b+a)
            case "*":
                stk = append(stk, b*a)
            case "-":
                stk = append(stk, b-a)
            case "/":
                stk = append(stk, b/a)
            }
        } else {
            // If the token is a number, convert it to integer and push it to the stack.
            num, _ := strconv.Atoi(token)
            stk = append(stk, num)
        }
    }
    // Return the final result from the stack.
    return stk[0]
}