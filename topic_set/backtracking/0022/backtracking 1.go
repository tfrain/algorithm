func generateParenthesis(n int) []string {
    var ret []string
    var backtrack func(int, int, []byte)
    backtrack = func(left, right int, combination []byte) {
        // If there are less opening brackets compared to closing ones, it's invalid and we terminate this branch
        if right < left {
            return
        }
        // If either left or right goes negative, we terminate this branch
        if left < 0 || right < 0 {
            return
        }
        // Every time we insert an opening (or closing) bracket, we decrement our opening (or closing) bracket count
        // If we've used up all our brackets, then we're done
        // We append the combination to the result array
        if left == 0 && right == 0 {
            ret = append(ret, string(combination))
            return
        }
        // We place an opening bracket, if there are any left to place. And recurse
        backtrack(left-1, right, append(combination, '('))
        // Then we place a closing bracket, if it won't break the syntax, and recurse
        backtrack(left, right-1, append(combination, ')'))
    }
    // We start the recursion with all opening and closing brackets
    backtrack(n, n, nil)
    return ret
}