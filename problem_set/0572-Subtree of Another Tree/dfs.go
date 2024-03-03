// function checks whether t is a subtree of s
func isSubtree(s, t *TreeNode) bool {
    if s == nil {
        // if s is null, return false
        return false
    }
    // if s and t are the same tree, return true
    if isSame(s, t) {
        return true
    }
    // else, recursively check if t is a subtree of left or right child of s
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// function checks whether two trees are identical
func isSame(t1, t2 *TreeNode) bool {
    if t1 == nil || t2 == nil {
        // if any of the trees is null, they are only identical if both are null
        return t1 == t2
    }
    // if the values of root nodes are not identical, the trees are not identical
    if t1.Val != t2.Val {
        return false
    }
    // else, recursively check if left and right child of both trees are identical
    return isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}