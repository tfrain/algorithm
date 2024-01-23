func buildTree(preorder []int, inorder []int) *TreeNode {
    // If the preorder slice is empty, return nil as the tree is empty
    if len(preorder) == 0 {
        return nil
    }
    
    // The root node value is always the first element in the preorder traversal
    node := preorder[0]
    
    // Find the index of the root in the inorder slice
    idx := indexOfInorder(inorder, node)
    
    // Recursively construct the left and right sub-trees
    root := &TreeNode {
        Val: node,
        Left: buildTree(preorder[1:idx+1], inorder[:idx]),
        Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
    }
    
    // Return the constructed tree
    return root
}

// Helper function to find the index of a node in the inorder slice
func indexOfInorder(inorder []int, node int) int{
    for i, v := range inorder {
        if v == node {
            return i
        }
    }
    return -1
}