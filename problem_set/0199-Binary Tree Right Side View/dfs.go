// rightSideView function returns the right side view of a binary tree.
// It uses a depth-first search approach to traverse the tree.
func rightSideView(root *TreeNode) []int {
    // values slice will store the values of the nodes from the right side view.
    var values []int

    // f is a recursive function that traverses the tree.
    var f func(*TreeNode, int)
    f = func(root *TreeNode, depth int) {
        // If the node is nil, return.
        if root == nil {
            return
        }

        // If the current depth is equal to the length of the values slice,
        // it means we are at a new level that we haven't visited before,
        // so we append the current node's value to the values slice.
        if depth == len(values) {
            values = append(values, root.Val)
        } else {
            // If we have already visited this level, we update the value
            // at this level with the current node's value.
            values[depth] = root.Val
        }

        // We first traverse the left subtree, then the right subtree.
        f(root.Left, depth+1)
        f(root.Right, depth+1)
    }
    // We start the traversal with the root at depth 0.
    f(root, 0)

    // We return the values from the right side view.
    return values
}


// rightSideView function returns the right side view of a binary tree.
func rightSideView(root *TreeNode) []int {
    // res slice will store the values of the nodes from the right side view.
    var res []int

    // bf is a recursive function that traverses the tree.
    var f func(root *TreeNode, level int)
    f = func(root *TreeNode, level int) {
        // If the node is nil, return.
        if root == nil {
            return
        }
        // If the current level is equal to the length of the res slice,
        // it means we are at a new level that we haven't visited before,
        // so we append the current node's value to the res slice.
        if level == len(res) {
            res = append(res, root.Val)
        }
        // We first traverse the right subtree, then the left subtree.
        f(root.Right, level+1)
        f(root.Left, level+1)
    }
    // We start the traversal with the root at level 0.
    f(root, 0)

    // We return the values from the right side view.
    return res
}