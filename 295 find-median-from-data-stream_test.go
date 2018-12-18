package leetcode

import (
	"fmt"
	"testing"
)

func insert2FM(mf *MedianFinder, slice []int) {
	for i, _ := range slice {
		mf.AddNum(slice[i])
	}
}

func TestFindMedian(t *testing.T) {
	mf := MConstructor()
	testArr := []int{2, 3}
	insert2FM(&mf, testArr)
	fmt.Printf("%f", mf.FindMedian())
}
