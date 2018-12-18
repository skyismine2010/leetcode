package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func crossMLoop(M int, head *ListNode) (*ListNode, *ListNode) {
	var prev, cur *ListNode
	cur = head
	prev = nil
	for i := 0; i < M-1; i++ {
		prev = cur
		cur = cur.Next
	}

	return prev, cur
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	loop := n - m
	if loop == 0 {
		return head
	}

	oldPrev, begin := crossMLoop(m, head)

	cur := begin
	next := cur.Next
	cur.Next = nil
	for i := 0; i < loop; i++ {
		prev := cur
		cur = next
		next = cur.Next
		cur.Next = prev
	}
	if oldPrev != nil {
		oldPrev.Next = cur
	}
	begin.Next = next
	if m != 1 {
		return head
	} else {
		return cur
	}
}
