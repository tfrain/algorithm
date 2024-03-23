func preorderTraversal(root *TreeNode) []int {
    // Initialize an empty stack
    stack := make([]*TreeNode, 0)
    
    // Initialize an empty slice to store the result
    var res []int

    // Loop until the stack is empty and the root is nil
    for len(stack) != 0 || root != nil {
        // If the root is nil, pop the last node from the stack and set it as the root
        if root == nil {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        
        // Append the value of the root to the result slice
        res = append(res, root.Val)
        
        // If the right child of the root is not nil, push it to the stack
        if root.Right != nil {
            stack = append(stack, root.Right)
        }
        
        // Move to the left child of the root
        root = root.Left
    }
    
    // Return the result slice
    return res
}