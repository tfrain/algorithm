
// Defines subtreeWithAllDeepest function
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {

    // Define a helper function 'find' which returns the subtree of the deepest level and its depth
    var find func(root *TreeNode) (*TreeNode, int)
    find = func(root *TreeNode) (*TreeNode, int) {

        // Base case: if root is nil, returns nil as node and 0 as depth
        if root == nil {
            return nil, 0
        }

        // Recursive case: find the deepest node on the left by calling find function recursively
        leftNode, leftDepth := find(root.Left)

        // Recursive case: find the deepest node on the right by calling find function recursively
        rightNode, rightDepth := find(root.Right)
        
        // Compare depths of left and right nodes; return the node with greater depth
        if leftDepth < rightDepth {
            return rightNode, rightDepth + 1
        } else if rightDepth < leftDepth {
            return leftNode, leftDepth + 1
        } else { // If the depths of both nodes are equal, return the root node
            return root, leftDepth + 1
        }
    }

    // Call the helper 'find' function on the root node
    node, _ := find(root)

    // Return the smallest subtree enclosed all the deepest nodes
    return node
}