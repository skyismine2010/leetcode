package leetcode

func farestIdx(a []int) int {
	scope := a[0]
	ret := 0
	for i, _ := range a {
		if a[i]+i > scope {
			scope = a[i] + i
			ret = i
		}
	}
	return ret
}

func canJump(nums []int) bool {
	i := 0
	for i < len(nums)-1 {
		scope := nums[i]
		if i+scope >= len(nums)-1 {
			return true
		}
		maxIdx := farestIdx(nums[i : i+scope+1])
		if maxIdx == 0 {
			return false
		}
		i += maxIdx
	}
	return true
}
