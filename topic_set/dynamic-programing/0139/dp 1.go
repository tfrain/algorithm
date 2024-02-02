func wordBreak(s string, wordDict []string) bool {
    n := len(s)

    // Create a memo array to store previous results
    memo := make([]int, n)

    // Declare recursive function
    var dp func(start int) bool
    dp = func(start int) bool {
        // If we've reached the end of the string, we return true
        if start >= n {
            return true
        }

        // If we've already calculated the result, retrieve it from the memo
        if memo[start] != 0 {
            return memo[start] == 1
        }

        // Loop through all the words in the dictionary
        for _, w := range wordDict {
            wLen := len(w)

            // If current word is longer than the remaining string, skip it
            if wLen+start > n {
                continue
            }

            // If current word does not match the string slice, skip it
            if w != s[start: start+wLen] {
                continue
            }

            // Recursively call the function
            // If we successfully partition the rest of the sentence, it means we can partition the whole sentence
            if dp(start+wLen) {
                // Memoize the result
                memo[start] = 1
                return true
            }
        }

        // If no word matches, we cannot partition the sentence so we memoize the result and return false
        memo[start] = -1
        return false
    }

    // Start the recursion from the beginning of the string
    return dp(0)
}

func wordBreak(s string, wordDict []string) bool {
    // Create a dp array, initialized true at 0
    dp := make([]bool, len(s)+1)
    dp[0] = true

    // Iterate through the dp array
    for i := 0; i < len(dp); i++ {
        // If current position is reachable
        if dp[i] {
            // Try to extend the partition by the length of each word in the dictionary
            for _, word := range wordDict {
                // If the word matches and the new position is within the bounds
                if i+len(word) <= len(s) && s[i:i+len(word)] == word {
                    // Mark the new position as reachable
                    dp[i+len(word)] = true
                }
            }
        }
    }

    // Check if we can reach the end of the string
    return dp[len(s)]
}