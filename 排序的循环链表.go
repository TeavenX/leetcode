package main

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insert(aNode *Node, x int) *Node {
	node := &Node{Val: x}
	if aNode == nil {
		node.Next = node
		return node
	}
	if aNode.Next == aNode {
		aNode.Next = node
		node.Next = aNode
		return aNode
	}
	cur := aNode
	for cur.Next != aNode {
		if cur.Val <= x && cur.Next.Val >= x {
			break
		}
		if cur.Val > cur.Next.Val && cur.Val >= x && cur.Next.Val >= x {
			break
		}
		if cur.Val > cur.Next.Val && cur.Val <= x && cur.Next.Val <= x {
			break
		}
		cur = cur.Next
	}
	node.Next = cur.Next
	cur.Next = node
	return aNode
}
