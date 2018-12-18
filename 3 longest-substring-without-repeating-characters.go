package leetcode

func mmax(a []int) int {
	val := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > val {
			val = a[i]
		}
	}
	return val

}

func isAlreadyExist(c string, charSet map[string]int) (bool, int) {
	if pos, ok := charSet[c]; ok {
		return true, pos
	}
	return false, 0
}

func flashSet(begin int, end int, origin string, charSet map[string]int) {
	for ; begin <= end; begin++ {
		delete(charSet, string(origin[begin]))
	}
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	charSet := make(map[string]int, 128)

	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	dp := make([]int, sLen)
	dp[0] = 1
	for right, _ := range s {
		if right != 0 {
			exist, pos := isAlreadyExist(string(s[right:right+1]), charSet)
			if exist == false {
				dp[right] = dp[right-1] + 1
			} else {
				flashSet(left, pos, s, charSet)
				dp[right] = len(charSet) + 1
				left = pos + 1
			}
		}

		charSet[s[right:right+1]] = right

	}

	return mmax(dp)
}
