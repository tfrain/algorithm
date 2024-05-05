func merge(intervals [][]int) [][]int {
    // Initialize an empty 2D slice of integers
    var res [][]int
    
    // Sort the intervals slice in ascending order based on the first element of each sub-slice
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    // Append the first sub-slice of intervals to res
    res = append(res, intervals[0])
    
    // Define a function max that returns the maximum of two integers
    max := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    // Iterate over the intervals slice starting from the second sub-slice
    for i := 1; i < len(intervals); i++ {
        // Define curr as the current sub-slice of intervals
        curr := intervals[i]
        
        // Define last as the last sub-slice of res
        last := res[len(res)-1]
        
        // If the first element of curr is less than or equal to the second element of last,
        // merge curr into last and update the second element of last to be the maximum of the second elements of last and curr
        if curr[0] <= last[1] {
            last[1] = max(last[1], curr[1])
        } else {
            // If the first element of curr is greater than the second element of last, append curr to res
            res = append(res, curr)
        }
    }

    // Return the merged intervals
    return res
}