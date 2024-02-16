// findMedianSortedArrays function returns the median value of combined nums1 and nums2
func findMedianSortedArrays(nums1, nums2 []int) float64 {
    // Ensuring nums1 is always the shorter array, swapping if necessary
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    // half is the sum of lengths of the two arrays divided by 2 (rounded up)
    half := (len(nums1) + len(nums2) + 1) / 2 

    // Declare variable to store the output median
    var median float64

    // Initializing binary search in nums1
    lo, hi := 0, len(nums1)

    // Start of binary search loop
    for lo <= hi {
        i := lo + (hi-lo)/2 // Finding partition for nums1

        j := half - i // Calculate partition for nums2 based on i

        // The following conditions evaluate our partition points i & j and adjust as needed
        if i > 0 && nums1[i-1] > nums2[j] {
            hi = i - 1 // move partition point in nums1 to left
        } else if i < len(nums1) && nums2[j-1] > nums1[i] {
            lo = i + 1 // move partition point in nums1 to right
        } else {
            // Partition point is at correct position, now we calculate maxLeft
            var maxLeft int
            if i == 0 {
                maxLeft = nums2[j-1]
            } else if j == 0 {
                maxLeft = nums1[i-1]
            } else {
                maxLeft = max(nums1[i-1], nums2[j-1])
            }

            // If the total length of both arrays is odd then the median is maxLeft
            if (len(nums1)+len(nums2)) % 2 == 1 {
                median = float64(maxLeft)
                break
            }

            // If the total length of both arrays is even then calculate minRight and find average of maxLeft and minRight
            var minRight int
            if i == len(nums1) {
                minRight = nums2[j]
            } else if j == len(nums2) {
                minRight = nums1[i]
            } else {
                minRight = min(nums1[i], nums2[j])
            }

            // Median is the average of maxLeft and minRight
            median = float64(maxLeft+minRight) / 2
            break
        }
    }

    // Return the calculated Median
    return median
}