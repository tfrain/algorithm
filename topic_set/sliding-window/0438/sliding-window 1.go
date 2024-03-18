func findAnagrams(s string, p string) []int {
    need, window := make(map[byte]int), make(map[byte]int)
    for i := range p {
        need[p[i]]++
    }
    left, right, valid := 0, 0, 0
    // record result
    res := []int{}
    for right < len(s) {
        c := s[right]
        right++
        // make some updates in window
        if _, ok := need[c]; ok {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        // check if the right window need to shrink
        for right-left >= len(p) {
            // update the result when a valid window is found
            if valid == len(need) {
                res = append(res, left)
            }
            d := s[left]
            left++
            // make some updates in window
            if _, ok := need[d]; ok {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    return res
}