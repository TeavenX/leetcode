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
func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	level := 0
	maxSum := 0
	maxLevel := 0
	queue = append(queue, root)
	for len(queue) > 0 {
		level++
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}
	}
	if maxLevel == 0 {
		maxLevel = level
	}
	return maxLevel
}
