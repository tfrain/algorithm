func longestConsecutive(nums []int) int {
    m := make(map[int]struct{})  // Initializes the map
    for _, v := range nums {
        m[v] = struct{}{}  // Inserts elements into the map
    }
    ret := 0
    for num := range m {  
        if _, ok := m[num-1]; ok {  // Checks if num-1 is already in the map
            continue
        }
        currNum := num  // Initializes the current number 
        currLen := 1     // Initializes the current length
        _, ok := m[currNum+1]  // Checks if currNum+1 is in the map
        for ok {  
            currLen++  
            currNum = currNum+1  // Increment current number for the next check
            _, ok = m[currNum+1] // Find if incremented number is in the map
        }
        if currLen > ret { // If current length is greater than the max length found so far
            ret = currLen  // Current length is the new max length
        }
    }
    return ret
}