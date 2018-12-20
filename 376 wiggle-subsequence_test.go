package leetcode

import "testing"

func TestWiggleMaxLength(t *testing.T) {
	if wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}) != 6 {
		t.Errorf("error")
	}

	if wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}) != 7 {
		t.Errorf("error")
	}

	if wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}) != 2 {
		t.Errorf("error")
	}

	if wiggleMaxLength([]int{0, 0}) != 1 {
		t.Errorf("error")
	}
}
