package leetcode

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	charMap := make(map[string]int)
//	for i := 0; i < len(s); i++ {
//		if s[i] == t[i] {
//			continue
//		}
//		if _, ok := charMap[string(s[i])+string(t[i])]; !ok {
//			charMap[string(t[i])+string(s[i])] += 1
//		} else {
//			charMap[string(s[i])+string(t[i])] -= 1
//			if charMap[string(s[i])+string(t[i])] == 0 {
//				delete(charMap, string(string(s[i])+string(t[i])))
//			}
//		}
//	}
//
//	if len(charMap) != 0 {
//		return false
//	}
//
//	return true
//}
//
//func main() {
//	s := "anagram"
//	t := "nagaram"
//	println(isAnagram(s, t))
//
//	s = "aacc"
//	t = "ccac"
//	println(isAnagram(s, t))
//
//	s = "dgqztusjuu"
//	t = "dqugjzutsu"
//	println(isAnagram(s, t))
//
//}

func isAnagram(s string, t string) bool {
	charMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		charMap[string(s[i])] += 1
	}

	for j := 0; j < len(t); j++ {
		charMap[string(t[j])] -= 1
		if charMap[string(t[j])] == 0 {
			delete(charMap, string(t[j]))
		}
	}

	return len(charMap) == 0

}
