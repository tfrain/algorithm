func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0]) // m is total rows and n is total columns in the matrix
    i, j := 0, n-1 // Initialize pointers. i pointing at the first row and j at the last column
    // Looping through the matrix diagonally starting from the top right cell until out of matrix boundaries
    for i < m && j >= 0 {
        // If target is greater than the current cell value move down the row
        if matrix[i][j] > target {
            j--
        // If target is lesser than current cell value move left the column
        } else if matrix[i][j] < target {
            i++
        // If target equals the current cell, return true as we found the target
        } else {
            return true
        }
    }
    // If loop is exhausted, target was not found in the matrix
    return false
}