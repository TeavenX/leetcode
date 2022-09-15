package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 5}
	fmt.Println(longestUnivaluePath(root))
}

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	maxL := 0
	var dfs func(node *TreeNode, pre int) int
	dfs = func(node *TreeNode, pre int) int {
		if node == nil {
			return 0
		}
		leftCount := dfs(node.Left, node.Val)
		rightCount := dfs(node.Right, node.Val)
		if leftCount+rightCount > maxL {
			maxL = leftCount + rightCount
		}
		if node.Val == pre {
			return max(leftCount, rightCount) + 1
		}
		return 0
	}
	if root == nil {
		return 0
	}
	dfs(root, root.Val+1)
	return maxL
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
