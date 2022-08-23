package main

import "strconv"

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func printTree(root *TreeNode) [][]string {
	maxDepth := 1
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if depth > maxDepth {
			maxDepth = depth
		}
		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
	}
	dfs(root, 1)
	result := make([][]string, maxDepth)
	for i := 0; i < maxDepth; i++ {
		result[i] = make([]string, 1<<maxDepth-1)
	}
	var dfsFill func(node *TreeNode, row, column int)
	dfsFill = func(node *TreeNode, row, column int) {
		result[row][column] = strconv.Itoa(node.Val)
		if node.Left != nil {
			dfsFill(node.Left, row+1, column-1<<(maxDepth-1-row-1))
		}
		if node.Right != nil {
			dfsFill(node.Right, row+1, column+1<<(maxDepth-1-row-1))
		}
	}
	dfsFill(root, 0, 1<<(maxDepth-1)-1)
	return result
}
