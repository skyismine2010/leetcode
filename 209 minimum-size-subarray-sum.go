package leetcode

func total(nums []int, l, r int) int {
	total := nums[l]
	for i := l + 1; i <= r; i++ {
		total += nums[i]
	}
	return total
}

func minSubArrayLen(s int, nums []int) int {
	l := 0
	r := 0

	current := 0
	length := len(nums)
	ret := 0

	if len(nums) == 0 {
		return 0
	}

	for l < length-1 {
		current = total(nums, l, r)
		if current >= s {
			if (r-l)+1 < ret || ret == 0 {
				ret = (r - l) + 1
			}
			l++
		} else {
			if r < length-1 {
				r++
			} else {
				l++
			}
		}
	}

	return ret
}
