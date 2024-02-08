// Function: findDuplicate (Binary Search Method)
// This function uses binary search on a sorted range of numbers to find a duplicate in an array.
func findDuplicate(nums []int) int {
    
    // We start the binary search with the smallest and the largest number.
    start, end := 1, len(nums)
    
    // We continue until the range becomes empty.
    for start <= end {
        
        // We find the middle number in the current range.
        mid := start + (end-start)/2
        
        // We count the numbers less than or equal to the middle number.
        cnt := 0 
        for _, val := range nums {
            if val <= mid {
                cnt++
            }
        }
        
        // If there are more numbers less than or equal to the middle number than the middle number itself, 
        // the duplicate is in the left half of the range.
        if cnt > mid {
            end = mid-1    
        } else { // Otherwise, the duplicate is in the right half of the range.
            start = mid+1 
        }
    }

    // The duplicate number is returned.
    return start
}