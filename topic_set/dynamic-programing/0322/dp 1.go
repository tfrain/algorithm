// Function to calculate the fewest number of coins that can make a specified amount
func coinChange(coins []int, amount int) int {
    // We create a special array 'amountMemo' where each position, i,
    // is initialized with amount + 1
    amountMemo := make([]int, amount+1)
    for i := range amountMemo {
        amountMemo[i] = amount + 1
    }

    // We set the minimum number of coins to make the amount of 0 is 0
    amountMemo[0] = 0

    // A helper function to find the minimum between two integers.
    min := func(a, b int) int {
        if a < b {
            return a
        }
        return b
    }

    // Iteratively build the 'amountMemo' array
    for i := range amountMemo {
        // For each coin, check if it can be involved in making the amount
        for _, v := range coins {
            // Skip when the current coin's value is greater than the current amount i,
            if i == 0 || i-v < 0 {
                continue
            }

            // Update the amountMemo[i] to the minimum between its current value
            // and 1 plus amountMemo[i-v] 
            amountMemo[i] = min(amountMemo[i], amountMemo[i-v]+1)
        }
    }

    // If we didn't find a combination to get the 'amount'
    if amountMemo[amount] == amount+1 {
        return -1
    }

    // Return the fewest number of coins that make a specified amount
    return amountMemo[amount]
}

// Function to calculate the fewest number of coins that you need to make up a specified amount
func coinChange(coins []int, amount int) int {
    // Initialize a memoization table with length = amount + 1
    memo := make([]int, amount+1)

    // A helper function to find the minimum between two integers
    min := func(a, b int) int {
        if a > b {
            return b
        }
        return a
    }

    // Recursive function to solve subproblems
    var dp func(amount int) int
    dp = func(amount int) int {
        // Base case: When amount is zero, 0 coins are required
        if amount == 0 {
            return 0
        }

        // Return the memoized result if it exists
        if memo[amount] != 0 {
            return memo[amount]
        }

        res := math.MaxInt32
        // Loop through each coin
        for _, coin := range coins {
            // Skip if current coin value is bigger than current amount
            if amount-coin < 0 {
                continue
            }

            // Calculate subproblem and skip if not possible
            subNum := dp(amount - coin)
            if subNum == -1 {
                continue
            }

            // Store the minimum value
            res = min(res, subNum+1)
        }

        // Update the memoization table
        if res == math.MaxInt32 {
            memo[amount] = -1
            return -1
        }

        memo[amount] = res
        return memo[amount]
    }

    // Call the recursive function initially with the provided amount
    return dp(amount)
}