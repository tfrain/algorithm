    func inorderTraversal(root *TreeNode) []int {
        // Initialize an empty slice to store the result
        var ret []int
        // Define a recursive function to traverse the tree
        var traverse func(root *TreeNode)
        traverse = func(root *TreeNode) {
            // If the node is nil, return
            if root == nil {
                return
            }
            // Traverse the left subtree
            traverse(root.Left)
            // Append the value of the current node to the result
            ret = append(ret, root.Val)
            // Traverse the right subtree
            traverse(root.Right)
        }
        // Start the traversal from the root
        traverse(root)
        // Return the result
        return ret
    }


    func inorderTraversal(root *TreeNode) []int {
        // Define a recursive function to traverse the tree and return the result
        var traverse func(root *TreeNode) []int
        traverse = func(root *TreeNode) []int {
            // Initialize an empty slice to store the result
            var ret []int
            // If the node is nil, return the empty slice
            if root == nil {
                return ret
            }
            // Append the result of traversing the left subtree
            ret = append(ret, traverse(root.Left)...)
            // Append the value of the current node
            ret = append(ret, root.Val)
            // Append the result of traversing the right subtree
            ret = append(ret, traverse(root.Right)...)
            // Return the result
            return ret
        }
        // Start the traversal from the root and return the result
        return traverse(root)
    }
