package main

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		lastLevel := len(queue)
		for i := 0; i < lastLevel-1; i++ {
			queue[i].Next = queue[i+1]
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if queue[lastLevel-1].Left != nil {
			queue = append(queue, queue[lastLevel-1].Left)
		}
		if queue[lastLevel-1].Right != nil {
			queue = append(queue, queue[lastLevel-1].Right)
		}
		queue = queue[lastLevel:]
	}
	return root
}

func connectV2(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		dummy := &Node{Val: -1}
		pre := dummy
		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}
	return root
}
