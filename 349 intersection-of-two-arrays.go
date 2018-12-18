package leetcode

func intersection(nums1 []int, nums2 []int) []int {
	setMap := make(map[int]int)
	retSetMap := make(map[int]int)
	var ret []int
	for i := 0; i < len(nums1); i++ {
		if _, ok := setMap[nums1[i]]; !ok {
			setMap[nums1[i]] = 0
		}
	}

	for i := 0; i < len(nums2); i++ {
		if _, ok := setMap[nums2[i]]; ok {
			retSetMap[nums2[i]] = 0
		}
	}

	for k, _ := range retSetMap {
		ret = append(ret, k)
	}
	return ret
}
