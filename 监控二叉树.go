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

const (
	NOCOVER = iota
	SETCAMERA
	COVERED
)

func minCameraCover(root *TreeNode) int {
	result := 0

	var postOrder func(*TreeNode) int
	postOrder = func(root *TreeNode) int {
		if root == nil {
			return COVERED
		}
		left := postOrder(root.Left)
		right := postOrder(root.Right)

		if left == COVERED && right == COVERED {
			return NOCOVER
		}

		if left == NOCOVER || right == NOCOVER {
			result++
			return SETCAMERA
		}

		if left == SETCAMERA || right == SETCAMERA {
			return COVERED
		}
		return -1
	}
	rootVal := postOrder(root)
	if rootVal == NOCOVER {
		result++
	}
	return result
}
