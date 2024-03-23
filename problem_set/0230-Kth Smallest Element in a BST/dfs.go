func kthSmallest(root *TreeNode, k int) int {
    var inOrder func(root *TreeNode, k int)
    cnt := 0
    val := 0
    // Define recursive in-order traversal.
    inOrder = func(root *TreeNode, k int) {
        if root ==  nil {
            return
        }
        // Visit left subtree.
        inOrder(root.Left, k)
        // Visit current node.
        cnt++
        // If 'k' smallest is found, save result and return.
        if cnt == k {
            val = root.Val
            return
        }
        // Visit right subtree.
        inOrder(root.Right, k)
    }
    inOrder(root, k)
    return val
}

func kthSmallest(root *TreeNode, k int) int {
    // Initialize an empty stack.
    stack := []*TreeNode{}
    for root != nil || len(stack) > 0 {
        // Insert all the left nodes in the stack.
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // At this point the first node in stack is the kth smallest, 
        // and it's not visited (the left subtree is implicitely visited without using stack).
        node := stack[len(stack)-1] // use the last element
        stack = stack[:len(stack)-1] // delete the last element
        k--
        // We found kth smallest, return its value.
        if k == 0 {
            return node.Val
        }
        // Now we need to process the right subtree.
        if node.Right != nil {
            root = node.Right
        }
    }
    return -1
}