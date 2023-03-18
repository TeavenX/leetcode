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
func reverseKGroup(head *ListNode, k int) *ListNode {
	temp := head
	for i := 0; i < k; i++ {
		if temp == nil {
			return head
		}
		temp = temp.Next
	}
	head, tail := reverse(head, k)
	tail.Next = reverseKGroup(temp, k)
	return head
}

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
	var pre *ListNode
	tail := head
	for k > 0 {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
		k--
	}
	return pre, tail
}
