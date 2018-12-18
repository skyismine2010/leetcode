package leetcode

func buildKey(str string) *[255]int {
	var retVal [255]int

	for i := 0; i < len(str); i++ {
		retVal[str[i]] += 1
	}
	return &retVal
}

func groupAnagrams(strs []string) [][]string {
	var retList [][]string
	wordMap := make(map[[255]int][]string)

	for i := 0; i < len(strs); i++ {
		key := buildKey(strs[i])
		wordMap[*key] = append(wordMap[*key], strs[i])
	}

	for _, v := range wordMap {
		retList = append(retList, v)
	}

	return retList
}
