package leetcode

import "testing"

func TestCanJump(t *testing.T) {
	if canJump([]int{2, 3, 1, 1, 4}) != true {
		t.Errorf("error")
	}

	if canJump([]int{3, 2, 1, 0, 4}) != false {
		t.Errorf("error")
	}

	if canJump([]int{0}) != true {
		t.Errorf("error")
	}

	if canJump([]int{2, 0, 0}) != true {
		t.Errorf("error")
	}

	if canJump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}) != true {
		t.Errorf("error")
	}

}
