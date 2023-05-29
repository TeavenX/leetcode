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
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	cache := make(map[int]bool)
	for _, node := range to_delete {
		cache[node] = true
	}
	ans := make([]*TreeNode, 0)
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		dl := dfs(node.Left)
		if dl {
			node.Left = nil
		}
		dl = dfs(node.Right)
		if dl {
			node.Right = nil
		}
		if cache[node.Val] {
			if node.Left != nil {
				ans = append(ans, node.Left)
			}
			if node.Right != nil {
				ans = append(ans, node.Right)
			}
			return true
		}
		return false
	}
	dl := dfs(root)
	if !dl {
		ans = append(ans, root)
	}
	return ans
}
