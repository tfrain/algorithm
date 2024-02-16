// Function hIndex calculates the H-index from an array of citations.
func hIndex(citations []int) int {
    // n is the number of papers the scientist has
    n := len(citations)
    
    // cnt is an array used to count the frequency of each citation number
    cnt := make([]int, n+1)
    
    // Iterate through the citations array
    for _, val := range citations {
    // If a citation count is larger than the number of papers,
    // we treat it as if it's equal to the number of papers
        if val > n {
            val = n
        }
        // Increment the count of citations for each value in the array
        cnt[val]++
    }
    
    // Initialize H-index as 0
    var h int
    
    // Try to find the maximum h such that the scientist has at least h papers 
    // that have been cited at least h times
    for i := n; i > 0; i-- {
    // If we find such a h, we break the loop
        if i <= cnt[i] {
            h = i
            break
        }
        // Else, we keep adding the count of the current index to the previous one 
        // to get a cumulative count
        cnt[i-1] += cnt[i]
    }
    // Return the found H-index
    return h
}