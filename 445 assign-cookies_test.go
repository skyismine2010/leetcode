package leetcode

import "testing"

func TestFindContentChildren(t *testing.T) {
	if findContentChildren([]int{1, 2, 3}, []int{1, 1}) != 1 {
		t.Errorf("error")
	}
	if findContentChildren([]int{1, 2}, []int{1, 2, 3}) != 2 {
		t.Errorf("error")
	}

	if findContentChildren([]int{1, 9}, []int{2, 1, 3, 6, 9, 7}) != 2 {
		t.Errorf("error")
	}

}
