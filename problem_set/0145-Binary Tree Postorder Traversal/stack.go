func postorderTraversal(root *TreeNode) []int {
    // Check if root is nil
    if root == nil {
        return nil
    }
    
    // Initialize return slice and stack
    var ret []int
    stack := []*TreeNode{}
    var prev *TreeNode
    
    // Loop until root is nil and stack is empty
    for root != nil || len(stack) > 0 {
        // Traverse to the leftmost node
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        
        // Get the top node from the stack
        node := stack[len(stack)-1]
        
        // Check if the right child of the node is nil or has been visited
        if node.Right == nil || node.Right == prev {
            // Pop the node from the stack
            stack = stack[:len(stack)-1]
            // Add the node value to the return slice
            ret = append(ret, node.Val)
            // Mark the node as visited
            prev = node
        } else {
            // Move to the right child
            root = node.Right
        }
    }
    
    // Return the postorder traversal
    return ret
}