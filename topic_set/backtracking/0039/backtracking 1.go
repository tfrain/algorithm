// combinationSum function takes an array of integers (candidates) and a target integer
// and returns all unique combinations in candidates where the candidate numbers sum to target.
func combinationSum(candidates []int, target int) [][]int {
    // Initialize an empty 2D slice to store the result
    res := [][]int{}

    // If the candidates slice is empty, return the empty result
    if len(candidates) == 0 {
        return res
    }

    // Declare a recursive function (backtrack) to find all combinations
    var backtrack func(candidates []int, start, sum int, track []int)
    backtrack = func(candidates []int, start, sum int, track []int) {
        // If the sum of the current combination equals the target
        if sum == target {
            // Make a copy of the current combination and add it to the result
            n := make([]int, len(track))
            copy(n, track)
            res = append(res, n)
            return
        }

        // If the sum of the current combination is greater than the target, return
        if sum > target {
            return
        }

        // Iterate over the candidates starting from the current index
        for i := start; i < len(candidates); i++ {
            // Call the backtrack function with the current candidate added to the sum and the combination
            backtrack(candidates, i, sum+candidates[i], append(track, candidates[i]))
        }
    }

    // Call the backtrack function with the initial parameters
    backtrack(candidates, 0, 0, nil)

    // Return the result
    return res
}

// combinationSum function takes an array of integers (candidates) and a target integer
// and returns all unique combinations in candidates where the candidate numbers sum to target.
func combinationSum(candidates []int, target int) [][]int {
    // Initialize an empty 2D slice to store the result
    var res [][]int

    // Get the length of the candidates slice
    n := len(candidates)

    // Declare a recursive function (backtrack) to find all combinations
    var backtrack func(start, sum int, subs []int)
    backtrack = func(start, sum int, subs []int) {
        // If the start index is out of range or the sum is less than 0, return
        if start >= n || sum < 0 {
            return
        }

        // If the sum equals 0 (target), add the current combination to the result
        if sum == 0 {
            tmp := make([]int, len(subs))
            copy(tmp, subs)
            res = append(res, tmp)
        }

        // Iterate over the candidates starting from the current index
        for i := start; i < n; i++ {
            // Call the backtrack function with the current candidate subtracted from the sum and added to the combination
            backtrack(i, sum-candidates[i], append(subs, candidates[i]))
        }
    }

    // Call the backtrack function with the initial parameters
    backtrack(0, target, nil)

    // Return the result
    return res
}