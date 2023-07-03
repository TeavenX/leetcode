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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)
	ans := &ListNode{}
	cur := ans
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Val = carry % 10
		carry /= 10
		if l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	for carry > 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = carry % 10
		carry /= 10
	}
	return reverse(ans)
}

func reverse(l *ListNode) *ListNode {
	var pre *ListNode = nil
	cur := l
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
