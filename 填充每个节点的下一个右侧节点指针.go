package main

func main() {
	a := Node{Val: 1}
	a.Left = &Node{Val: 2}
	a.Right = &Node{Val: 3}
	a.Left.Left = &Node{Val: 4}
	a.Left.Right = &Node{Val: 5}
	a.Right.Left = &Node{Val: 6}
	a.Right.Right = &Node{Val: 7}
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/*
初始化一个队列，将当前这一层遍历到的节点按`左节点`、`右节点`的顺序放入队列
每一层遍历完，都从队列里取出节点，设置当前节点的下一个节点为队列的下一个节点
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	nodeCount := 0
	queue = append(queue, root)
	nodeCount++
	var previous *Node
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
		nodeCount--
		if previous != nil {
			previous.Next = cur
		}
		previous = cur
		if nodeCount == 0 {
			cur.Next = nil
			previous = nil
			nodeCount = len(queue)
		}
	}
	return root
}

func connectV2(root *Node) *Node {
	if root == nil {
		return root
	}
	left, right := root.Left, root.Right
	for left != nil {
		left.Next = right
		right = right.Left
		left = left.Right
	}
	connectV2(root.Left)
	connectV2(root.Right)
	return root
}
