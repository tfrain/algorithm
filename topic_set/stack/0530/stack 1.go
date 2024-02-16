// getMinimumDifference calculates the minimum absolute difference between the values of any two nodes.
func getMinimumDifference(root *TreeNode) int {
    // Initialize minimum difference with maximum possible integer value.
    min := math.MaxInt64
    // Initialize a stack to help in the iteration.
    stack := []*TreeNode{}
    // Define previous node pointer for tracking the last visited node.
    var prev *TreeNode

    // Iterate until all nodes are traversed.
    for root != nil || len(stack) > 0 {
        // Keep adding the left subtree nodes into the stack until we reach the leftmost node.
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }

        // Now process the top node. First, pop the top element from the stack.
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // If there is a previously visited node, check if the difference with the current node is smaller than current minimum difference.
        if prev != nil && root.Val-prev.Val < min {
            min = root.Val - prev.Val
        }
        // Update the previous node pointer to current node.
        prev = root
        
        // Move to the right subtree.
        root = root.Right
    }

    // Return the minimum difference found.
    return min
}