func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		path += strconv.Itoa(node.Val)

		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			return
		}

		if node.Left != nil {
			dfs(node.Left, path+"->")
		}
		if node.Right != nil {
			dfs(node.Right, path+"->")
		}
	}

	dfs(root, "")
	return paths
}
