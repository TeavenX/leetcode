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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	var tail *ListNode
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		val := v1 + v2 + carry
		val, carry = val%10, val/10
		if head == nil {
			head = &ListNode{Val: val}
			tail = head
		} else {
			tail.Next = &ListNode{Val: val}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	ans := &ListNode{}
	var pre *ListNode = nil
	cur := ans
	for l1 != nil && l2 != nil {
		v := l1.Val + l2.Val + carry
		cur.Val = v % 10
		carry = v / 10
		l1 = l1.Next
		l2 = l2.Next
		cur.Next = &ListNode{}
		pre = cur
		cur = cur.Next
	}
	for l1 != nil {
		v := l1.Val + carry
		cur.Val = v % 10
		carry = v / 10
		l1 = l1.Next
		cur.Next = &ListNode{}
		pre = cur
		cur = cur.Next
	}
	for l2 != nil {
		v := l2.Val + carry
		cur.Val = v % 10
		carry = v / 10
		l2 = l2.Next
		cur.Next = &ListNode{}
		pre = cur
		cur = cur.Next
	}
	for carry > 0 {
		cur.Val = carry % 10
		carry /= 10
		cur.Next = &ListNode{}
		pre = cur
		cur = cur.Next
	}
	pre.Next = nil
	return ans
}
