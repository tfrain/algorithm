func judgePoint24(cards []int) bool {
    var helper func([]float64) bool
    helper = func(nums []float64) bool {
        // Base case: If only one number left, check if it's approximately 24
        if len(nums) == 1 {
            return math.Abs(nums[0]-24) < 1e-6
        }
        for i := range nums {
            for j := range nums {
                // Only consider pair where i < j
                if i == j {
                    continue
                }
                var nxt []float64
                // Store remaining numbers to process in next iteration
                for k, num := range nums {
                    if k != i && k != j {
                        nxt = append(nxt, num)
                    }
                }
                // All six potential results from each pair
                for _, num := range getNums(nums[i], nums[j]) {
                    // If one way of computing the numbers is close to 24, return true
                    if helper(append(nxt, num)) {
                        return true
                    }
                }
            }
        }
        return false
    }
    nums := make([]float64, len(cards))
    for i, v := range cards {
        nums[i] = float64(v)
    }
    return helper(nums)
}

func getNums(a, b float64) []float64 {
    // 6 possible combinations of a and b using basic arithmetics
    return []float64{a + b, a - b, b - a, a * b, a / b, b / a}
}