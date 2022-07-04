package main

import "math"

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
func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, lower, upper int) bool
	dfs = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		return dfs(node.Left, lower, node.Val) && dfs(node.Right, node.Val, upper)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

func isValidBSTV2(root *TreeNode) bool {
	cache := make([]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		cache = append(cache, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	for i := 1; i < len(cache); i++ {
		if cache[i] <= cache[i-1] {
			return false
		}
	}
	return true
}
