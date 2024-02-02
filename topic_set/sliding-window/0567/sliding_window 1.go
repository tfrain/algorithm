func checkInclusion(s1 string, s2 string) bool {
    // Create two maps for characters in s1 (need) and the current window (window) respectively
    need, window := make(map[byte]int), make(map[byte]int)
    // Store the frequency of characters in s1 to the need map
    for i := range s1 {
        need[s1[i]]++
    }
    // Initialize left, right pointers to represent the window, and a float for the valid count
    left, right, valid := 0, 0, 0
    // Start moving the right pointer
    for right < len(s2) {
        c := s2[right]
        right++
        // Increase the count of the character in the window map and if it matches the requirement, increase valid count
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        // Keep the window size to the length of s1
        while right-left >= len(s1) {
            // If all characters have met the requirement, return true
            if valid == len(need) {
                return true
            }
            // Start to remove the left character from the window
            d := s2[left]
            left++
            // Decrease the count in the window and if it was valid before, decrease valid count
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    // If no valid permutation is found, return false
    return false
}