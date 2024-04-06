
func flatten(root *TreeNode)  {
    // Declare a function variable preorder
    var preorder func(root *TreeNode) *TreeNode

    // Define the preorder function
    preorder = func(root *TreeNode) *TreeNode {
        // If the root is nil, return nil
        if root == nil {
            return nil
        }

        // Recursively call the preorder function on the left child of the root
        left := preorder(root.Left)

        // Recursively call the preorder function on the right child of the root
        right := preorder(root.Right)

        // Set the left child of the root to nil
        root.Left = nil

        // Set the right child of the root to the left child
        root.Right = left

        // Initialize a temporary variable to the root
        tmp := root

        // Traverse the right children of the root until the right child is nil
        for tmp.Right != nil {
            tmp = tmp.Right
        }

        // Set the right child of the last node to the right child of the root
        tmp.Right = right

        // Return the root
        return root
    }

    // Call the preorder function on the root
    preorder(root)
}