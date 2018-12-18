package leetcode

import (
	"strconv"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func isVaildIP(s string) bool {
	size := len(s)
	intVal, _ := strconv.Atoi(s)

	if s[0] == '0' && size != 1 {
		return false
	}

	if intVal > 255 {
		return false
	}

	return true
}

func splitIPStrByDot(s string, N int) []string {
	var resultList []string
	if len(s) > 3*N || len(s) < 1*N {
		return nil
	}

	if N == 1 {
		var tmp []string
		if isVaildIP(s) {
			return append(tmp, s)
		}
		return nil
	}

	for i := 0; i < Min(3, len(s)); i++ {
		prefix := s[:i+1]
		if isVaildIP(prefix) == false {
			continue
		}

		ret := splitIPStrByDot(s[i+1:], N-1)
		if ret == nil {
			continue
		}
		size := len(ret)
		for j := 0; j < size; j++ {
			resultList = append(resultList, prefix+"."+ret[j])
		}
	}
	return resultList

}

func restoreIpAddresses(s string) []string {
	ret := splitIPStrByDot(s, 4)
	return ret
}
