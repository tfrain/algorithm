func lengthOfLongestSubstring(s string) int {
    res := 0
    // Create an index dictionary
    index := make(map[byte]int)
    for i, j := 0, 0; j < len(s); j++ {
        // If character already exists in dictionary, move the window to next index of duplicate
        if index[s[j]] > i {
            i = index[s[j]]
        }
        // If the current substring is longer, update the longest substring length
        if j-i+1 > res {
            res = j - i + 1
        }
        // Update the current character's index in the dictionary
        index[s[j]] = j + 1
    }
    return res
}