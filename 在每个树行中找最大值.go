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
func largestValues(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		maxNum := queue[0].Val
		for i := 0; i < size; i++ {
			if maxNum < queue[i].Val {
				maxNum = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		result = append(result, maxNum)
		queue = queue[size:]
	}
	return
}

func largestValuesDFS(root *TreeNode) (result []int) {
	if root == nil {
		return
	}
	var dfs func(node *TreeNode, height int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if height == len(result) {
			result = append(result, node.Val)
		} else {
			result[height] = max(result[height], node.Val)
		}
		dfs(node.Left, height+1)
		dfs(node.Right, height+1)
	}
	dfs(root, 0)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
