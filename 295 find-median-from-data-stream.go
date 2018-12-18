package leetcode

import (
	"container/heap"
	. "github.com/skyismine2010/leetcode/MyStruct"
)

type MedianFinder struct {
	leftHeap  MaxHeap
	rightHeap MinHeap
	size      int
}

/** initialize your data structure here. */
func MConstructor() MedianFinder {
	m := MedianFinder{nil, nil, 0}
	return m
}

func (this *MedianFinder) adjustHeaps() {
	leftLen := len(this.leftHeap)
	rightLen := len(this.rightHeap)
	if leftLen > rightLen && leftLen-rightLen > 1 {
		leftMax := heap.Pop(&this.leftHeap)
		heap.Push(&this.rightHeap, leftMax)
	}

	if leftLen < rightLen && rightLen-leftLen > 1 {
		rightMin := heap.Pop(&this.rightHeap)
		heap.Push(&this.leftHeap, rightMin)
	}
}

func (this *MedianFinder) AddNum(num int) {
	leftMax := this.leftHeap.Top()
	if num <= leftMax {
		heap.Push(&this.leftHeap, num)
	} else {
		heap.Push(&this.rightHeap, num)
	}

	this.adjustHeaps()
}

func (this *MedianFinder) FindMedian() float64 {
	leftMax := this.leftHeap.Top()
	rightMin := this.rightHeap.Top()
	if len(this.leftHeap) == len(this.rightHeap) {
		return float64(float64(leftMax+rightMin) / 2)
	}
	if len(this.leftHeap) > len(this.rightHeap) {
		return float64(leftMax)
	} else {
		return float64(rightMin)
	}

}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
