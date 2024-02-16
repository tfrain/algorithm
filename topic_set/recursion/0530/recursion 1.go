// getMinimumDifference calculates the minimum absolute difference between the values of any two nodes.
func getMinimumDifference(root *TreeNode) int {
    // Initialize minimum difference with maximum possible integer value.
    min := math.MaxInt64

    // Define previous node pointer for tracking the last visited node.
    var prev *TreeNode
    // Define recursive function (f) to in-order traverse the tree.
    var f func(*TreeNode)
    f = func(root *TreeNode) {
        // Return if the node is null.
        if root == nil {
            return
        }

        // Traverse the left subtree.
        f(root.Left)
        // If there is a previously visited node, check if the difference with the current node is smaller than current minimum difference.
        if prev != nil && root.Val-prev.Val < min {
            min = root.Val - prev.Val
        }
        // Update the previous node pointer to current node.
        prev = root
        // Traverse the right subtree.
        f(root.Right)
    }
    // Initiate the in-order traversal with root node.
    f(root)

    // Return the minimum difference found.
    return min
}