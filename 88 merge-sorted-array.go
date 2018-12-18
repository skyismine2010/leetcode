package leetcode

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	for n > 0 && m > 0 {
		if nums2[n-1] > nums1[m-1] {
			nums1[pos] = nums2[n-1]
			n--
		} else {
			nums1[pos] = nums1[m-1]
			m--
		}
		pos--
	}
	for n > 0 {
		nums1[pos] = nums2[n-1]
		pos--
		n--
	}

	fmt.Println(nums1)
}
