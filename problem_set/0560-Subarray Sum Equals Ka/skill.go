func subarraySum(nums []int, k int) int {
    prefix := map[int]int{0: 1} // Hash map stores prefix sum and its occurrence count
    var sum, cnt int
    for _, num := range nums{
        sum += num // `sum` holds the sum of `nums[:i]`
        cnt += prefix[sum-k] // `cnt` records number of cases where subarray sum is `k`
        prefix[sum]++ // Increment prefix sum by 1
    }
    return cnt
}