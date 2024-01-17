package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion && bottom-up && post-order traversal && DFS
func binaryTreePaths0(root *TreeNode) []string {
	if root == nil {
		return nil
	}
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

type StackNode struct {
	Node *TreeNode
	Path string
}

// Recursion && up-bottom && Preorder traversal && DFS
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

// Stack && up-bottom && Preorder traversal && DFS
func binaryTreePaths1(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var paths []string
	stack := []StackNode{{Node: root, Path: strconv.Itoa(root.Val)}}
	var b strings.Builder
	for len(stack) > 0 {
		stackNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := stackNode.Node
		path := stackNode.Path

		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
		}

		if node.Right != nil {
			b.Reset()
			b.WriteString(path)
			b.WriteString("->")
			b.WriteString(strconv.Itoa(node.Right.Val))
			stack = append(stack, StackNode{Node: node.Right, Path: b.String()})
		}
		if node.Left != nil {
			b.Reset()
			b.WriteString(path)
			b.WriteString("->")
			b.WriteString(strconv.Itoa(node.Left.Val))
			stack = append(stack, StackNode{Node: node.Left, Path: b.String()})
		}
	}

	return paths
}

// Stack && bottom-up && Preorder traversal && DFS
func binaryTreePathsBottomUpPreOrder(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	var stack, resultStack []StackNode
	stack = append(stack, StackNode{Node: root, Path: strconv.Itoa(root.Val)})

	// Pre-order traversal but store in stack instead of processing immediately
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If leaf node, add to result stack for later processing
		if node.Node.Left == nil && node.Node.Right == nil {
			resultStack = append(resultStack, node)
		}

		// Push children to stack for pre-order traversal
		if node.Node.Right != nil {
			newPath := node.Path + "->" + strconv.Itoa(node.Node.Right.Val)
			stack = append(stack, StackNode{Node: node.Node.Right, Path: newPath})
		}
		if node.Node.Left != nil {
			newPath := node.Path + "->" + strconv.Itoa(node.Node.Left.Val)
			stack = append(stack, StackNode{Node: node.Node.Left, Path: newPath})
		}
	}

	// Process resultStack for bottom-up construction
	var paths []string
	for i := len(resultStack) - 1; i >= 0; i-- {
		paths = append(paths, resultStack[i].Path)
	}

	return paths
}
