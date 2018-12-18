package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func deleteNode(prev *ListNode, cur *ListNode) *ListNode {
	next := cur.Next
	prev.Next = next
	cur.Next = nil
	return next
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	next := cur.Next
	val := cur.Val
	for next != nil {
		if val == next.Val {
			next = deleteNode(cur, next)
		} else {
			val = next.Val
			cur = next
			next = next.Next
		}
	}
	return head
}
