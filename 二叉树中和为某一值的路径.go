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
func pathSum(root *TreeNode, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && sum+node.Val == target {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs(node.Left, sum+node.Val)
		dfs(node.Right, sum+node.Val)
	}
	dfs(root, 0)
	return ans
}
