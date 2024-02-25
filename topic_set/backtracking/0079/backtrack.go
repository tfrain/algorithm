func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])   // Compute the Dimensions of the board
    res := false  // Initialize the result as false
    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}  // Pre-defined directions for traversing: right, down, left, up.

    // DFS coupled with Backtracking
    var dfs func(start, i, j int)
    dfs = func(start, i, j int) {
        if start == len(word) {   // If the word is fully found, update the result as true, and return.
            res = true
            return
        }
        // If the index is out of bounds or characters don't match, return immediately.
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[start] {
            return
        }
        tmp := board[i][j]  // Save the original value of the cell.
        board[i][j] = '1'   // Mark the current cell as visited.

        // Traverse in all four directions recursively.
        for _, d := range directions {
            dfs(start+1, i+d[0], j+d[1])   
        }

        // Important step: Backtrack by restoring the original value of the cell after exploring all paths.
        board[i][j] = tmp  
    }

    // Start DFS from every cell.
    for i := range board {
        for j := range board[i] {
            dfs(0, i, j)
        }
    }
    return res   // Return the final result.
}