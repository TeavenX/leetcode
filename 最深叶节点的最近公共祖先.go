package main

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
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	depth := 0
	ans := root
	var dfs func(cur int, node *TreeNode) int
	dfs = func(cur int, node *TreeNode) int {
		if node == nil {
			depth = max(depth, cur)
			return cur
		}
		l := dfs(cur+1, node.Left)
		r := dfs(cur+1, node.Right)
		if l == r && l == depth {
			ans = node
		}
		return max(l, r)
	}
	dfs(0, root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
