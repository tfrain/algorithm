func canFinish(numCourses int, prerequisites [][]int) bool {
    // Create adjacency list graph from edges and compute in-degrees
    startGraph := make([][]int, numCourses)
    inDegree := make([]int, numCourses)
    for _, p := range prerequisites {
        startGraph[p[1]] = append(startGraph[p[1]], p[0])
        inDegree[p[0]]++
    }	
    
    // Find nodes with in-degree = 0
    inDegreeZeroList := make([]int, 0)
    for node, d := range inDegree {
        if d == 0 {
            inDegreeZeroList = append(inDegreeZeroList, node)
        }
    }
    
    // Iteratively remove nodes with in-degree = 0 and update in-degrees
    for i := 0; i < len(inDegreeZeroList); i++ {
        for _, j := range startGraph[inDegreeZeroList[i]] {
            inDegree[j]--
            // Add the node to zero in-degree list if it becomes 0
            if inDegree[j] == 0 {
                inDegreeZeroList = append(inDegreeZeroList, j)
            }
        }
    }
    
    // Check if all nodes are in zero in-degree list
    return len(inDegreeZeroList) == numCourses
}