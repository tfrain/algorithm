package binarytreepaths

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	// base case
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	paths := append(binaryTreePaths(root.Left), binaryTreePaths(root.Right)...)
	v := strconv.Itoa(root.Val)
	var b strings.Builder
	for i, path := range paths {
		b.Reset()
		b.WriteString(v)
		b.WriteString("->")
		b.WriteString(path)

		paths[i] = b.String()
	}
	return paths
}
