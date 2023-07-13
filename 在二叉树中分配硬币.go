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
func distributeCoins(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := dfs(node.Left), dfs(node.Right)
		ans += abs(l) + abs(r)
		return node.Val + l + r - 1
	}
	dfs(root)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
