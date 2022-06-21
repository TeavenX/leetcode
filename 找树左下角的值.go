package main

import "fmt"

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(findBottomLeftValue(root))
	fmt.Println(findBottomLeftValueBFS(root))
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
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var dfs func(node *TreeNode)
	depth := 0
	maxDepth := 0
	maxNum := root.Val
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		depth++
		if depth > maxDepth {
			maxDepth = depth
			maxNum = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
		depth--
	}
	dfs(root)
	return maxNum
}

func findBottomLeftValueBFS(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := root.Val
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if i == 0 {
				result = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}
	return result
}
