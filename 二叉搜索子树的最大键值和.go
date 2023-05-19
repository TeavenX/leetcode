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

var (
	INF = int(4e8)
)

type ReturnVal struct {
	min, max, sum int
	isBST         bool
}

func maxSumBST(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) ReturnVal
	dfs = func(node *TreeNode) ReturnVal {
		if node == nil {
			return ReturnVal{INF, -INF, 0, true}
		}
		lv, rv := dfs(node.Left), dfs(node.Right)
		if lv.isBST && rv.isBST && lv.max < node.Val && node.Val < rv.min {
			ans = max(ans, lv.sum+rv.sum+node.Val)
			return ReturnVal{min(lv.min, node.Val), max(rv.max, node.Val), lv.sum + rv.sum + node.Val, true}
		} else {
			return ReturnVal{0, 0, 0, false}
		}
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
