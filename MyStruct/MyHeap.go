package MyStruct

import "math"

type MinHeap []int

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *MinHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *MinHeap) Top() int {
	if len(*h) == 0 {
		return math.MaxInt32
	} else {
		return (*h)[0]
	}
}

type MaxHeap []int

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *MaxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *MaxHeap) Top() int {
	if len(*h) == 0 {
		return math.MinInt32
	} else {
		return (*h)[0]
	}
}
