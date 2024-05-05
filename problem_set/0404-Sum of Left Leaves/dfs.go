func sumOfLeftLeaves(root *TreeNode) int {
    // If the root is nil, return 0
    if root == nil {
        return 0
    }
    var sum int
    // Check if the left child of the root is a leaf node
    if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
        // If it is, add its value to the sum
        sum += root.Left.Val
    } else {
        // If it's not, recursively call the function on the left child
        sum += sumOfLeftLeaves(root.Left)
    }
    // Recursively call the function on the right child, as it might have left leaf nodes
    return sum + sumOfLeftLeaves(root.Right)
}                                          