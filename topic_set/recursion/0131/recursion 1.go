func partition(s string) [][]string {
    var ret [][]string  // Final result array

    if len(s) == 0 {  // If the string is empty: base case
        return [][]string{{}}  // Return an empty string as an element of the slice of slices
    }
    if len(s) == 1 {  // If the string comprise of a single character: base case
        ret = append(ret, []string{s})  // Return this string as an element of the slice
        return ret
    }
    // For each substring
    for i := 0; i < len(s); i++ {
        if sub := s[:i+1]; checkPalindome(sub, 0, len(sub)-1) {  // If the substring ('sub') is palindrome
            subs := partition(s[i+1:])  // Recursive call to 'partition' for the rest of the string after 'i+1'- to get all possible partitions
            // Traverse each partition from 'subs'
            for _, v := range subs {
                // Append 'sub' before each partition
                partition := append([]string{sub}, v...)
                ret = append(ret, partition)  // Append the new partition to the 'ret'
            }
        }
    }
    return ret  // Return the final result array
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