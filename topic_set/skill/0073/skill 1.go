func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0]) // Get matrix dimensions
    zeroIdx := make([][2]int, 0) // Initialize zero index holder
    // Identify zero indices
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                zeroIdx = append(zeroIdx, [2]int{i, j})
            }
        }
    }
    // Return if there are no zeroes
    if len(zeroIdx) == 0 {
        return
    }
    // In-place function to set zeroes
    setZero := func(row, col int) {
        for i := 0; i < m; i++ {
            matrix[i][col] = 0
        }
        for i := 0; i < n; i++ {
            matrix[row][i] = 0
        }
    }
    
    // Set zeroes using identified indices
    for _, idx := range zeroIdx {
        setZero(idx[0], idx[1])
    }
}

func setZeroes(matrix [][]int) {
    m, n := len(matrix), len(matrix[0]) // Get matrix dimensions
    row, col := make([]bool, m), make([]bool, n) // Initialize boolean arrays
    // Record zero-indexing
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                row[i] = true
                col[j] = true
            }
        }
    }
    // Set rows to zeroes
    for i, v := range row {
        if v {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }
    // Set columns to zeroes
    for i, v := range col {
        if v {
            for j := 0; j < m; j++ {
                matrix[j][i] = 0
            }
        }
    }
}