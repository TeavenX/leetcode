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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	cache := make(map[*ListNode]bool)
	for head != nil {
		if exist := cache[head]; exist {
			return true
		}
		cache[head] = true
		head = head.Next
	}
	return false
}

func hasCycleV2(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		if fast == slow {
			return true
		}
		slow = slow.Next
	}
	return false
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle20220527(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	pre := &ListNode{Next: head}
	slow, fast := pre.Next, pre.Next
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			if fast == slow {
				return true
			}
		}
	}
	return false
}
