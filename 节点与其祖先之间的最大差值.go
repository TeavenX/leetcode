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
func maxAncestorDiff(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, l, h int)
	dfs = func(node *TreeNode, l, h int) {
		if node == nil {
			return
		}
		if l == -1 {
			l = node.Val
		}
		if h == -1 {
			h = node.Val
		}
		ans = max(ans, max(abs(l-node.Val), abs(h-node.Val)))
		l = min(l, node.Val)
		h = max(h, node.Val)
		dfs(node.Left, l, h)
		dfs(node.Right, l, h)
	}
	dfs(root, -1, -1)
	return ans
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
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
