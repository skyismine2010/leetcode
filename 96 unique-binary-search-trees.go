package leetcode

func getNums(num int, dp *[]int) int {
	ret := 0
	if num < len(*dp) {
		return (*dp)[num]
	}
	for i := 0; i < num; i++ {
		leftN := getNums(i, dp)
		rightN := getNums(num-i-1, dp)
		ret += (leftN * rightN)
	}
	*dp = append(*dp, ret)

	return ret
}

func numTrees(n int) int {
	var dp []int
	dp = append(dp, 1)
	dp = append(dp, 1)
	dp = append(dp, 2)
	dp = append(dp, 5)

	if n < len(dp) {
		return dp[n]
	}

	for num := 4; num <= n; num++ {
		getNums(n, &dp)
	}
	return dp[n]
}
