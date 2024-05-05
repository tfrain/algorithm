func longestCommonPrefix(strs []string) string {
    // Return an empty string if the slice is empty or the first string in the slice is empty
    if len(strs) == 0 || len(strs[0]) == 0 {
        return ""
    }
    // Initialize minLen as the maximum value of an integer
    minLen := math.MaxInt32
    // Iterate over the strings in the slice
    for _, str := range strs {
        // If the length of a string is less than minLen, update minLen
        if len(str) < minLen {
            minLen = len(str)
        }
    }
    // Initialize idx as 0 and n as the length of strs
    idx, n := 0, len(strs)
    // Initialize ret as an empty slice of bytes
    ret := []byte{}
    // While idx is less than minLen
    for idx < minLen {
        // Define c as the idx-th byte of the first string in strs
        c := strs[0][idx]
        // Iterate over the strings in strs starting from the second
        for i := 1; i < n; i++ {
            // If the idx-th byte of a string is not equal to c, return ret as a string
            if c != strs[i][idx] {
                return string(ret)
            }
        }
        // Append c to ret
        ret = append(ret, c)
        // Increment idx by 1
        idx++
    }
    // Return ret as a string
    return string(ret)
}

func longestCommonPrefix1(strs []string) string {
    // Return an empty string if the slice is empty or the first string in the slice is empty
    if len(strs) == 0 || len(strs[0]) == 0 {
        return ""
    }
    // Initialize res as an empty slice of bytes
    var res []byte
    // Initialize idx as 0
    idx := 0
    // While idx is less than the length of the first string in strs
    for idx < len(strs[0]) {
        // Initialize a flag as true
        flag := true
        // Define c as the idx-th byte of the first string in strs
        c := strs[0][idx]
        // Iterate over the strings in strs starting from the second
        for _, str := range strs[1:] {
            // If the length of a string is less than or equal to idx or its idx-th byte is not equal to c, set flag to false and break the loop
            if len(str) <= idx || str[idx] != c {
                flag = false
                break
            }
        }
        // If flag is false, break the loop
        if !flag {
            break
        }
        // Append c to res
        res = append(res, c)
        // Increment idx by 1
        idx++
    }
    // Return res as a string
    return string(res)
}