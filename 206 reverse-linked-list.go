package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	next := cur.Next
	for next != nil {
		prev := cur
		cur = next
		next = cur.Next
		cur.Next = prev
	}

	return cur
}
