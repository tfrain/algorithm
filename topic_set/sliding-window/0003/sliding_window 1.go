func lengthOfLongestSubstring(s string) int {
    // Create a window dictionary
    window := make(map[byte]int)
    left, right, res := 0, 0, 0
    // Slide the window to the right
    for right < len(s) {
        r := s[right]
        window[r]++
        right++
        // If there are duplicates, slide the window to the left
        for window[r] > 1 {
            l := s[left]
            window[l]--
            left++
        }
        // Update the longest substring length
        if res < right-left {
            res = right-left
        }
    }
    return res
}