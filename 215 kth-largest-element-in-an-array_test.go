package leetcode

import (
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	if findKthLargest([]int{2, 1, 5, 7, 8, 4}, 5) != 2 {
		t.Errorf("error")
	}

	if findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) != 5 {
		t.Errorf("error")
	}

	if findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) != 4 {
		t.Errorf("error")
	}

}
