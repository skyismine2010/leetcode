package MyStruct

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedList(a []int) *ListNode {
	var cur, head *ListNode

	for i := 0; i < len(a); i++ {
		newNode := ListNode{a[i], nil}
		if cur != nil {
			cur.Next = &newNode
			cur = cur.Next
		} else {
			cur = &newNode
			head = cur
		}
	}
	return head
}

func PrintLinkedList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")

}
