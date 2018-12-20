package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	var i, j int
	var satisfied int
	sort.Ints(g)
	sort.Ints(s)

	for j < len(s) {
		if s[j] >= g[i] {
			satisfied++
			i++
			j++
			if i == len(g) {
				break
			}
		} else { //这个糖果没办法满足当前的孩子
			j++
		}
	}
	return satisfied
}
