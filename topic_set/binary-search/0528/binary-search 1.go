// Solution struct incorporates the preSum slice (prefix sums of weights)
type Solution struct {
    preSum []int
}

// Constructor function initializes the Solution struct with a
// preSum slice. It calculates the prefix sums of the weights in 'w' array
func Constructor(w []int) Solution {
    // Get the length of weights array
    n := len(w)
    // Create preSum array with size one more than 'n'
    preSum := make([]int, n+1)
    // Compute the prefix sums for 'w' and store in 'preSum'
    for i := 1; i <= n; i++ {
        preSum[i] = preSum[i-1] + w[i-1]
    }
    // Return a new Solution struct with the computed 'preSum'
    return Solution{
        preSum: preSum,
    }
}

// PickIndex function chooses an index randomly in accordance with the weights 
func (this *Solution) PickIndex() int {
    // Get the length of preSum array
    n := len(this.preSum)
    // Generate a random target weight in range 1 to total weight
    target := rand.Intn(this.preSum[n-1]) + 1
    // Initialize binary search bounds
    l, r := 0, n
    // Perform binary search to find the target in the 'preSum' array
    for l < r {
        m := l + (r-l)/2 // Compute mid-point
        // If the mid-point value is less than target, then search in the right half
        if this.preSum[m] < target {
            l = m + 1
        } else { // Else, search in the left half
            r = m
        }
    }
    // Return the selected index, after adjusting for 1-based to 0-based index conversion
    return l - 1
}