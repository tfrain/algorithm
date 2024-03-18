func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     // base case
     if root == nil {
         return nil
     }
     if root == p || root == q {
         return root
     }
     left := lowestCommonAncestor(root.Left, p, q)
     right := lowestCommonAncestor(root.Right, p, q)

     // root
     if left != nil && right != nil {
         return root
     }
     // another base case
     if left == nil && right == nil {
         return nil
     }
     // send to parent node
     if left != nil {
         return left
     }
     return right
}