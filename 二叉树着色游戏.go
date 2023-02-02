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
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	bcount := 0
	var dfs func(node *TreeNode)
	var redNode *TreeNode
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Val == x {
			redNode = node
			return
		}
		bcount++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	if bcount > n>>1 {
		return true
	}
	bcount = 0
	dfs(redNode.Left)
	if bcount > n>>1 {
		return true
	}
	bcount = 0
	dfs(redNode.Right)
	if bcount > n>>1 {
		return true
	}
	return false
}
