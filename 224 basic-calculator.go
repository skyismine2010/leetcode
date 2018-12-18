package leetcode

//
//import (
//	"container/list"
//	"strconv"
//	"strings"
//)
//
//const (
//	InitState = iota
//	NumState
//	OperState
//)
//
//func isNumBegin(s string) bool {
//	if s[0] >= '0' && s[0] <= '9' {
//		return true
//	}
//	return false
//}
//
//func calculate(s string) int {
//	var number int
//	s = strings.TrimSpace(s)
//	numStack := list.New()
//	OperStack := list.New()
//	computeFlag := false
//	i := 0
//	state := InitState
//	var l, r int
//
//	for i < len(s) {
//		switch state {
//		case InitState:
//			if isNumBegin(s[i:]) {
//				state = NumState
//				l = i
//				r = i
//
//			} else {
//				state = OperState
//			}
//			i++
//
//		case NumState:
//			if isNumBegin(s[i:]) {
//				r++
//				i++
//			} else {
//				number, _ = strconv.Atoi(s[l : r+1])
//				numStack.PushFront(number)
//				state = OperState
//				number = 0
//			}
//
//		case OperState:
//			if s[i] == '+' || s[i] == '-' {
//				OperStack.PushFront(s[i])
//				i++
//			} else if isNumBegin(s[i:]) {
//				state = NumState
//				l = i
//				r = i
//			}
//		}
//	}
//}
