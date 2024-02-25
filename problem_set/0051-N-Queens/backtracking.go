func solveNQueens(n int) [][]string {
    // Fields array will store the current state of the chessboard 
    fields := make([][]byte, n)

    // Loop to initialize fields array with '.'
    for i := range fields {
        fields[i] = make([]byte, n)
        for j := range fields[i] {
            fields[i][j] = '.'
        }
    }

    // Initialize the results array to store valid chessboard configurations
    var ret [][]string

    // Recursive function for backtracking
    var bt func(int)
    bt = func(row int) {
        // Base case: if all rows have been consumed
        if row == n {
            ret = append(ret, toString(fields))
            return
        }
        for i := 0; i < n; i++ {
            // If placing a Queen at this cell is invalid, skip this cell
            if !validN(fields, row, i) {
                continue
            }
            // Else, place a Queen and move on to the next row
            fields[row][i] = 'Q'
            bt(row + 1)
            // Backtrack: remove the Queen from this cell
            fields[row][i] = '.'
        }
    }

    // Start from the first row
    bt(0)

    // Return all valid configurations
    return ret
}

// Function to check if placing a Queen at cell (row, col) is valid
func validN(fields [][]byte, row, col int) bool {
    n := len(fields)
    // Check if this column has a Queen already
    for i := 0; i < row; i++ {
        if fields[i][col] == 'Q' {
            return false
        }
    }
    // Check if the diagonals have a Queen already
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if fields[i][j] == 'Q' {
            return false
        }
    }
    for i, j  := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if fields[i][j] == 'Q' {
            return false
        }
    }
    // Placement is valid
    return true
}

// Function to convert the current chessboard configuration to a string
func toString(fields [][]byte) (res []string) {
    for i := range fields {
        res = append(res, string(fields[i]))
    }
    return
}