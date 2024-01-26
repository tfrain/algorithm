func partitionLabels(s string) []int {
     // helper function to get max of two integers
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    // create a map to store the last index of each character in the given string
    lastIdx := make(map[byte]int)
    for idx := range s {
        lastIdx[s[idx]] = idx
    }
    
    start, end, ans := 0, 0, []int{}
    for idx := range s {
        // update the furthest 'end' index for current partition
        end = max(end, lastIdx[s[idx]])
        
        // if we've processed all characters up to current 'end', partition is found
        if idx == end {
            ans = append(ans, idx-start+1)
            // prepare for next partition
            start = idx+1
        }
    }
    
    // return the lengths of partitions
    return ans
}