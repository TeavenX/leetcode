package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesError(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for head.Next == head {
		head = head.Next
	}

	dummy := &ListNode{0, head}
	cur := dummy.Next
	pre := dummy
	for cur != nil {
		if cur.Val == cur.Next.Val {
			cur = cur.Next.Next
		} else {
			pre.Next = cur
			cur = cur.Next
			pre = pre.Next
		}
	}
	return dummy.Next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for head.Next == head {
		head = head.Next
	}

	dummy := &ListNode{0, head}
	pre, cur := dummy, head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if pre.Next == cur {
			// 上面的循环没有执行
			pre = pre.Next
		} else {
			pre.Next = cur.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}
