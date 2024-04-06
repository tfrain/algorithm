// intersection is a function that takes two integer slices as input
// and returns a slice that contains the intersection of the two input slices.
func intersection(nums1, nums2 []int) []int {
    // Create a map to store the elements of the first slice.
    m := make(map[int]struct{})
    // Iterate over the first slice and add each element to the map.
    for _, num := range nums1 {
        m[num] = struct{}{}
    }

    // Create a slice to store the intersection of the two slices.
    nums := make([]int, 0)
    // Iterate over the second slice.
    for _, num := range nums2 {
        // If the current element of the second slice is in the map,
        // it means it is also in the first slice, so add it to the intersection slice.
        if _, ok := m[num]; ok {
            nums = append(nums, num)
            // After adding the element to the intersection slice, remove it from the map
            // to avoid adding duplicate elements to the intersection slice.
            delete(m, num)
        }
    }
    // Return the intersection slice.
    return nums
}