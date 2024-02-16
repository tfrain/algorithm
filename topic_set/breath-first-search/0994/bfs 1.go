// Implementing a Breadth-first search (BFS) to solve the problem
func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])   // Getting lengths of rows and cols
    queue := [][2]int{}  
    for r := 0; r < m; r++ {
        for l := 0; l < n; l++ {
            if grid[r][l] == 2 {
                // record all rotten orange
                queue = append(queue, [2]int{r, l})    // Add rotten oranges to queue
            }
        }
    }
    directions := [][2]int{{1,0}, {-1, 0}, {0, 1}, {0, -1}}
    // All possible directions where the orange can rot

    minutes := 0
    for len(queue) > 0 {
        for _, location := range queue {
            queue = queue[1:]   // Move to next orange in the queue
            for _, dir := range directions {
                r, c := location[0]+dir[0], location[1]+dir[1]
                if r < 0 || r >= m || c < 0 || c >= n {
                    continue   // Skip if coordinates are out of grid or the orange is rotten already 
                }
                if grid[r][c] != 1 {
                    continue    // Skip if it's not a fresh orange
                }
                grid[r][c] = 2  // Make the fresh orange rot
                queue = append(queue, [2]int{r, c})   // Add the new rotten orange to the queue
            }
        }
        if len(queue) > 0 {
            minutes++    // Increase the time as all oranges in the current queue rotten
        }
    }
    // Check if there are still remaining fresh oranges
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] == 1 {
                return -1   // Return -1 if a fresh orange still exists
            }
        }
    }
    return minutes   // Minimal time for all oranges to rot
}