package leetcode

type node struct {
	v    int
	next *node
}

type MyStack struct {
	head *node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{nil, 0}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	newNode := &node{x, this.head}
	this.head = newNode
	this.size++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.head != nil {
		headVal := this.head.v
		this.head = this.head.next
		this.size--
		return headVal
	}

	return 0

}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.head.v
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}
