    func inorderTraversal(root *TreeNode) []int {
        // Initialize an empty slice to store the result
        var ret []int
        // Initialize an empty stack to store the nodes
        stack := []*TreeNode{}
        // While there are nodes to process
        for root != nil || len(stack) > 0 {
            // Push all the nodes on the left to the stack
            for root != nil {
                stack = append(stack, root)
                root = root.Left
            }
            // Pop the node from the top of the stack
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // Append the value of the node to the result
            ret = append(ret, node.Val)
            // Process the right subtree
            root = node.Right
        }
        // Return the result
        return ret
    }