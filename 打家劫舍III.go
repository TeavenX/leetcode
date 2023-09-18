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
func rob(root *TreeNode) int {
	cache := make(map[*TreeNode][2]int)
	var dfs func(node *TreeNode, pre int) int
	dfs = func(node *TreeNode, pre int) (res int) {
		if node == nil {
			return 0
		}
		if v := cache[node][pre]; v > 0 {
			return v
		}
		defer func() {
			t := cache[node]
			t[pre] = res
			cache[node] = t
		}()
		res = dfs(node.Left, 0) + dfs(node.Right, 0)
		if pre == 1 {
			return res
		}
		res = max(res, dfs(node.Left, 1)+dfs(node.Right, 1)+node.Val)
		return res
	}
	return max(dfs(root, 0), dfs(root, 1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
