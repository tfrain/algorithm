func minWindow(s string, t string) string {
    window, need := make(map[byte]int), make(map[byte]int)
    // Populate the need map with the count of each character in t
    for i := range t {
        need[t[i]]++
    }
    // Initialize pointers for the sliding window and a counter for valid characters
    left, right, valid := 0, 0, 0
    // Initialize the start and end of the minimum window
    sLeft, sRight := 0, len(s)+1

    // Extend the right edge of the window
    for right < len(s) {
        r := s[right]
        right++
        // If the character is needed, add it to the window and check validity
        if _, ok := need[r]; ok {
            window[r]++
            if window[r] == need[r] {
                valid++
            }
        }
        // Shrink the window from the left and update the minimum window
        for len(need) == valid {
            if sRight-sLeft > right-left {
                sRight = right
                sLeft = left
            }
            l := s[left]
            left++
            // Update window and valid count when removing characters
            if _, ok := need[l]; ok {
                if window[l] == need[l] {
                    valid--
                }
                window[l]--
            }
        }
    }
    // Return the minimum window substring or an empty string
    if sLeft == 0 && sRight == len(s)+1 {
        return ""
    }
    return s[sLeft:sRight]
}
