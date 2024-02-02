func longestPalindrome(s string) string {
    var ret string  // A string to store the result
    // Define a function to check if the substring is a palindrome
    palindrome := func(s string, l, r int) string {
        // Continue while the characters at l and r are equal and within the bounds of s
        for l >= 0 && r < len(s) && s[l] == s[r] { 
            l-- // Move towards start of string
            r++ // Move towards end of string
        }
        // Return the identified palindrome
        return s[l+1 : r]
    }

    // Iterate over the string
    for i := 0; i < len(s); i++ {
        first := palindrome(s, i, i)    // Check for odd length palindromes
        second := palindrome(s, i, i+1) // Check for even length palindromes
        if len(ret) < len(first) { // If the current palindrome is longer than the recorded one
            ret = first // Replace ret with the current palindrome
        }
        if len(ret) < len(second) { // If the current palindrome is longer than the recorded one
            ret = second // Replace ret with the current palindrome
        }
    }
    return ret  // Return the longest palindrome
}

func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""  // If input string is empty, return an empty string
    }
    // Function to find the minimum of two integers
    min := func(a, b int) int {  
        if a < b {
            return a
        }
        return b
    }
    // Preprocess the string to uniformly handle odd and even length palindromes
    newStr := "#" 
    for _, c := range s {
        newStr += string(c) + "#" // Insert special character between each characters
    }

    n := len(newStr)             // Length of the preprocessed string
    p := make([]int, n)          // To store the length of the palindrome centered at each character
    center, right := 0, 0        // Initialize the center and right boundary of the palindrome
    maxLen, start := 0, 0        // Initialize the maximum length and starting position of the longest palindrome

    // Iterate through the expanded string
    for i := 0; i < n; i++ { 
        if i < right {
            // mirror position for i with respect to center is 2*center - i
            // i-center=center-x, x=2*center-i
            // p[i] is minimum of what is already known and what is yet to be tested
            p[i] = min(p[2*center-i], right-i)
        } else {
            // If i is beyond current right boundary, initialize p[i] to 1
            p[i] = 1
        }

        // Attempt to expand the palindrome centered at i
        for i+p[i] < n && i-p[i] >= 0 && newStr[i+p[i]] == newStr[i-p[i]] {
            p[i]++  // Increase the length of palindrome as long as symmetry holds
        }

        // Update center and right if the palindrome centered at i expands beyond right
        if i+p[i]-1 > right {
            center = i
            right = i + p[i] - 1
        }

        // Update maximum length and starting position of longest palindrome if necessary
        if p[i]-1 > maxLen {
            maxLen = p[i] - 1
            // Convert position in new string back to position in original string
            start = (i - maxLen) / 2
        }
    }
    // Return the longest palindrome from the original string s
    return s[start : start+maxLen]  
}