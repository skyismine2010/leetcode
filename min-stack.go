package leetcode

//type listNode MyStruct {
//	v    int
//	next *listNode
//}
//
//type MyStack MyStruct {
//	head *listNode
//	size int
//}
//
///** Push element x onto stack. */
//func (this *MyStack) Push(x int) {
//	newNode := &listNode{x, this.head}
//	this.head = newNode
//	this.size++
//}
//
///** Removes the element on top of the stack and returns that element. */
//func (this *MyStack) Pop() int {
//	if this.head != nil {
//		headVal := this.head.v
//		this.head = this.head.next
//		this.size--
//		return headVal
//	}
//
//	return 0
//
//}
//
///** Get the top element. */
//func (this *MyStack) Top() int {
//	return this.head.v
//}
//
///** Returns whether the stack is empty. */
//func (this *MyStack) Empty() bool {
//	if this.size == 0 {
//		return true
//	}
//	return false
//}
//
//type MinStack MyStruct {
//	dataStack *MyStack
//	minStack  *MyStack
//}
//
///** initialize your data structure here. */
//func MiniStack() MinStack {
//	return MinStack{&MyStack{nil, 0}, &MyStack{nil, 0}}
//}
//
//func (this *MinStack) Push(x int) {
//
//	this.dataStack.Push(x)
//	if this.minStack.Empty() {
//		this.minStack.Push(x)
//	} else {
//		if this.minStack.head.v > x {
//			this.minStack.Push(x)
//		} else {
//			this.minStack.Push(this.minStack.head.v)
//		}
//	}
//
//}
//
//func (this *MinStack) Pop() {
//	this.dataStack.Pop()
//	this.minStack.Pop()
//}
//
//func (this *MinStack) Top() int {
//	return this.dataStack.head.v
//}
//
//func (this *MinStack) GetMin() int {
//	return this.minStack.head.v
//}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := MiniStack();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
