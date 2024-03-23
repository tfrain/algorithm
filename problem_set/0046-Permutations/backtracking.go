// permute function generates all possible permutations of the given integer slice.
func permute(nums []int) [][]int {
    // ret is used to store all the permutations.
    var ret [][]int

    // Declare a function backtrack that will be used to generate permutations.
    var backtrack func([]int, []bool)

    // Define the function backtrack.
    backtrack = func(one []int, used []bool) {
        // n is the length of the input slice.
        // nlen is the length of the current permutation.
        n, nlen := len(nums), len(one)

        // If the length of the current permutation is equal to the length of the input slice,
        // it means we have found a permutation.
        if nlen == n {
            // Add the found permutation to the result slice.
            ret = append(ret, one)
            return
        }

        // Iterate over the input slice.
        for i := 0; i < n; i++ {
            // If the current element is already used in the current permutation, skip it.
            if used[i] {
                continue
            }
            // Mark the current element as used.
            used[i] = true
            // Call backtrack recursively with the current element added to the current permutation.
            backtrack(append(one, nums[i]), used)
            // After the recursive call, mark the current element as unused.
            used[i] = false
        }
    }

    // Initialize a slice to keep track of the used elements.
    used := make([]bool, len(nums))
    // Call the backtrack function to start generating permutations.
    backtrack(nil, used)

    // Return the result slice containing all the permutations.
    return ret
}