// pathSum counts the number of paths in a binary tree that sum up to targetSum
func pathSum(root *TreeNode, targetSum int) int {
    // numPath is a helper function to recursively calculate the number of targetSum paths
    var numPath func(root *TreeNode, sum int, top bool) int
    numPath = func(root *TreeNode, sum int, top bool) int {
        // Base case: if the root is nil return 0
        if root == nil {
            return 0
        }
        var num int
        // If the current node's value equals sum, increment num
        if root.Val == sum {
            num += 1
        }
        // When top flag is true, recursively calculate the number of paths for the left and right children
        if top {
            num += numPath(root.Left, sum, true) + numPath(root.Right, sum, true)
        }
        // Decrease sum by the current node's value and recursively calculate the number of paths for the left and right leaves
        sum -= root.Val
        num += numPath(root.Left, sum, false) + numPath(root.Right, sum, false)
        
        return num
    }
    // Return the final number of paths equal to targetSum starting from root
    return numPath(root, targetSum, true)
}

// pathSum counts the number of paths in a binary tree that sum up to sum
func pathSum(root *TreeNode, sum int) int {
    var res int
    // path is a helper function to recursively calculate the number of sum paths
    var path func(root *TreeNode, sum int, includeRoot bool)
    path = func(root *TreeNode, sum int, includeRoot bool) {
        // Base case: if the root is nil return
        if root == nil {
            return
        }
        // If the current node's value equals sum, increment res
        if sum == root.Val {
            res++
        }
        // When includeRoot flag is true, recursively calculate the number of paths for the left and right children
        if includeRoot {
            path(root.Left, sum, true)
            path(root.Right, sum, true)
        }
        // Decrease sum by the current node's value and recursively calculate the number of paths for the left and right leaves
        path(root.Left, sum-root.Val, false)
        path(root.Right, sum-root.Val, false)
    }
    
    // Recursive call with root and sum
    path(root, sum, true)
    
    // Return the final number of paths equal to sum starting from root
    return res
}