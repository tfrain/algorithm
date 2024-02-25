func partition(s string) [][]string {
    var ret [][]string  // Final result array
    n := len(s)  // Length of the input string
    
    // Initialized the function for backtracking
    var bt func(subs []string, start int)
    bt = func(subs []string, start int) {
        if start == n {  // If start index is equal to the length of the string
            tmp := make([]string, len(subs))  // Clone the 'subs' slice into a temporary one
            copy(tmp, subs)
            ret = append(ret, tmp)  // append the temp slice to the final result
        }
        // For each substring
        for i := start; i < n; i++ {
            if checkPalindome(s, start, i) {  // Recursive call to function checkPalindrome() to check if the substring of s from 'start' to 'i' indexes is a palindrome
                bt(append(subs, s[start:i+1]), i+1)  // Call backtracking function recursively by appending the new substring to 'subs' and incrementing 'i' by 1
            }
        }
    }
    bt(nil, 0)  // Initial call to backtracking function, stating start=0 and 'subs' as nil
    return ret  // Return final result array
}

func checkPalindome(s string, l, r int) bool {
    for l < r {
        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }
    return true
}