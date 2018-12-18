package leetcode

import "container/heap"
import . "github.com/skyismine2010/leetcode/MyStruct"

func findKthLargest(nums []int, k int) int {
	h := make(MinHeap, k)
	copy(h, nums[:k])
	heap.Init(&h)

	for i, _ := range nums[k:] {
		if h.Top() < nums[i+k] {
			heap.Pop(&h)
			heap.Push(&h, nums[i+k])
		}
	}

	return h.Top()
}
