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
func goodNodes(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, mx int)
	dfs = func(node *TreeNode, mx int) {
		if node == nil {
			return
		}
		if node.Val >= mx {
			ans++
		}
		dfs(node.Left, max(node.Val, mx))
		dfs(node.Right, max(node.Val, mx))
	}
	dfs(root, -1e5)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
