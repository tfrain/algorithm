// maxProduct finds the maximum product of a subarray within an integer array.
func maxProduct(nums []int) int {
    // Initialize maxV, minV and res with the first element of the array.
    maxV, minV, res := nums[0], nums[0], nums[0]

    // Define a function to find the maximum of two integers.
    max := func(a, b int) int{
        if a > b {
            return a
        }
        return b
    }
    // Define a function to find the minimum of two integers.
    min := func(a, b int) int{
        if a < b {
            return a
        }
        return b
    }
    // Iterate over the array, starting from the second element.
    for _, v := range nums[1:] {
        // If the current element is negative, swap maxV and minV.
        if v < 0 {
            minV, maxV = maxV, minV
        }
        // Update minV and maxV by comparing the product of the current element and the previous minV/maxV with the current element.
        minV = min(minV*v, v)
        maxV = max(maxV*v, v)
        // If maxV is greater than res, update res.
        if maxV > res {
            res = maxV
        }
    }
    // Return the maximum product of a subarray.
    return res
}