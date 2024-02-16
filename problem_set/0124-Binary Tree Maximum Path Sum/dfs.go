func maxPathSum(root *TreeNode) int {
    ret := math.MinInt32  // Initialization of max sum with the smallest possible integer
    max := func(a, b int) int { // Auxiliary function to get maximum of two integers
        if a > b {
            return a
        }
        return b
    }
    var maxPath func(root *TreeNode) int // Declaring our recursive helper function
    maxPath = func(root *TreeNode) int {
        if root == nil {
            return 0   // Base case: if node is nil, contribute 0 to the sum
        }
        leftMax := max(maxPath(root.Left), 0) // Maximum path sum from left child
        rightMax := max(maxPath(root.Right), 0) // Maximum path sum from right child
        ret = max(ret, root.Val+leftMax+rightMax) // Update our maximum path sum
        return root.Val + max(leftMax, rightMax)  // Return max path through current node
    }
    maxPath(root) // Invoke recursive function with root node
    return ret // Return the maximum path sum
}