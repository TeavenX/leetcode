package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorderRecur(root *Node) []int {
	result := make([]int, 0)
	var preOrder func(root *Node)
	preOrder = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		for _, node := range root.Children {
			preOrder(node)
		}
	}
	preOrder(root)
	return result
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return result
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return result
}
