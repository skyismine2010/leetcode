package leetcode

import "strings"

func wordPattern(pattern string, str string) bool {
	pattenMap := make(map[string]string)
	strMap := make(map[string]string)

	strVec := strings.Split(str, " ")

	if len(strVec) != len(pattern) {
		return false
	}

	for idx, str := range strVec {
		pattenVal, pattenExist := pattenMap[string(pattern[idx])]
		strVal, strExist := strMap[str]
		if (!pattenExist && strExist) || (pattenExist && !strExist) {
			return false
		}

		if !pattenExist && !strExist {
			pattenMap[string(pattern[idx])] = str
			strMap[str] = string(pattern[idx])
		} else {
			if pattenVal != str {
				return false
			}

			if strVal != string(pattern[idx]) {
				return false
			}
		}

	}

	return true
}
