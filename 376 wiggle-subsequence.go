package leetcode

const (
	BeginState = iota
	UpState
	DownState
)

func wiggleMaxLength(nums []int) int {
	state := BeginState
	maxLen := 1

	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		switch state {
		case BeginState:
			if nums[i]-nums[i-1] > 0 {
				state = UpState
				maxLen++
			}
			if nums[i]-nums[i-1] < 0 {
				state = DownState
				maxLen++
			}
		case UpState:
			if nums[i]-nums[i-1] < 0 {
				state = DownState
				maxLen++
			}
		case DownState:
			if nums[i]-nums[i-1] > 0 {
				state = UpState
				maxLen++
			}
		}
	}
	return maxLen

}
