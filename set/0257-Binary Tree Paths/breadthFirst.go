package binarytreepaths

import "strconv"

type nodePath struct {
	node *TreeNode
	path string
}

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	queue := []nodePath{{node: root, path: ""}}

	for len(queue) > 0 {
		np := queue[0]
		queue = queue[1:]

		curNode := np.node
		curPath := np.path
		if curNode.Left == nil && curNode.Right == nil {
			paths = append(paths, curPath+strconv.Itoa(curNode.Val))
		}
		if curNode.Left != nil {
			queue = append(queue, nodePath{node: curNode.Left, path: curPath + strconv.Itoa(curNode.Val) + "->"})
		}
		if curNode.Right != nil {
			queue = append(queue, nodePath{node: curNode.Right, path: curPath + strconv.Itoa(curNode.Val) + "->"})
		}
	}
	return paths
}
