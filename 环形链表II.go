package main

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: -4}
	head.Next.Next.Next.Next = head.Next
	detectCycleV2(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := hasCycle(head)
	if node == nil {
		return nil
	}
	cycleCount := 1
	temp := node.Next
	for temp != node {
		cycleCount++
		temp = temp.Next
	}
	slow, fast := head, head
	for cycleCount > 0 {
		fast = fast.Next
		cycleCount--
	}
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func hasCycle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil && slow != fast {
		fast = fast.Next
		slow = slow.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	if fast == nil {
		return nil
	}
	return slow
}

func detectCycleV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node := hasCycle(head)
	if node == nil {
		return nil
	}
	slow := head
	node = node.Next
	for slow != node {
		slow = slow.Next
		node = node.Next
	}
	return slow
}
