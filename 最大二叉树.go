package main

func main() {

}

type Stack struct {
	data []*TreeNode
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Push(val *TreeNode) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() *TreeNode {
	if s.Len() == 0 {
		return nil
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

func (s *Stack) Top() *TreeNode {
	if s.Len() == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Bottom() *TreeNode {
	if s.Len() == 0 {
		return nil
	}
	return s.data[0]
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	stack := Stack{}
	for _, num := range nums {
		node := &TreeNode{Val: num}
		for stack.Len() > 0 && stack.Top().Val < num {
			node.Left = stack.Pop()
		}
		if stack.Len() > 0 {
			stack.Top().Right = node
		}
		stack.Push(node)
	}
	return stack.Bottom()
}

func constructMaximumBinaryTreeV2(nums []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		max := left
		for i := left; i <= right; i++ {
			if nums[max] < nums[i] {
				max = i
			}
		}
		node := &TreeNode{Val: nums[max]}
		node.Left = helper(left, max-1)
		node.Right = helper(max+1, right)
		return node
	}
	return helper(0, len(nums)-1)
}
