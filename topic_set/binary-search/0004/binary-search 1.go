// This function is used to find the median of two sorted arrays.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Calculate total length of the two arrays
    if l := len(nums1) + len(nums2); l%2 == 0 {    // Check if total length is even
        // If it is even, the median is average of l/2-th number and (l/2 - 1)-th number
        return (findKth(nums1, nums2, l/2-1) + findKth(nums1, nums2, l/2)) / 2.0
    } else { 
        // If it is odd, the median is the l/2-th number
        return findKth(nums1, nums2, l/2)
    }
}

// This function is used to find the K-th number in two sorted arrays
func findKth(nums1, nums2 []int, k int) float64 {
    // Loop until finding the k-th number
    for {
        l1, l2 := len(nums1), len(nums2) // Get lengths of nums1 and nums2
        m1, m2 := l1/2, l2/2 // Calculate the mid-indices of nums1 and nums2
        if l1 == 0 { // If nums1 is empty, the k-th number must be in nums2
            return float64(nums2[k])
        } else if l2 == 0 { // If nums2 is empty, the k-th number must be in nums1
            return float64(nums1[k])
        } else if k == 0 { // If k is zero, the k-th number is the smaller number of the fronts of nums1 and nums2
            if n1, n2 := nums1[0], nums2[0]; n1 > n2 {
                return float64(n2)
            } else {
                return float64(n1)
            }
        }

        // If k is not zero, split the problem into a smaller one
        if k <= m1+m2 {
            // If k is less than or equal to the sum of mid-indices, the k-th number must be in the front halves
            // So we cut the latter half of the longer array
            if nums1[m1] <= nums2[m2] {
                nums2 = nums2[:m2]
            } else {
                nums1 = nums1[:m1]
            }
        } else {
            // If k is larger than the sum of mid-indices, the k-th number must be in the back halves or the front half of the shorter array
            // So we cut the front half of the shorter array and update k
            if nums1[m1] <= nums2[m2] {
                nums1 = nums1[m1+1:]
                k -= m1 + 1
            } else {
                nums2 = nums2[m2+1:]
                k -= m2 + 1
            }
        }
    }
}