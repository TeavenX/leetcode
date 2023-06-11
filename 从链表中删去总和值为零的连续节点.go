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
func removeZeroSumSublists(head *ListNode) *ListNode {
	sentinel := &ListNode{0, head}
	sum := 0
	cur := sentinel
	cache := make(map[int]*ListNode)
	for cur != nil {
		sum += cur.Val
		cache[sum] = cur
		cur = cur.Next
	}
	sum = 0
	cur = sentinel
	for cur != nil {
		sum += cur.Val
		cur.Next = cache[sum].Next
		cur = cur.Next
	}
	return sentinel.Next
}
