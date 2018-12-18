package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

/**
 * Definition for a binary tree listNode.
 * type TreeNode MyStruct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func reverse(s []int) []int {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}

	return s
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	flag := 1
	var retList [][]int
	var currentNodeList []*TreeNode
	if root == nil {
		return retList
	}
	currentNodeList = append(currentNodeList, root)
	for len(currentNodeList) != 0 {
		var nextNodeList []*TreeNode
		var ret []int
		for i := 0; i < len(currentNodeList); i++ {
			ret = append(ret, currentNodeList[i].Val)
			if currentNodeList[i].Left != nil {
				nextNodeList = append(nextNodeList, currentNodeList[i].Left)
			}
			if currentNodeList[i].Right != nil {
				nextNodeList = append(nextNodeList, currentNodeList[i].Right)
			}
		}

		if len(ret) != 0 {
			if flag%2 == 0 {
				ret = reverse(ret)
			}
			retList = append(retList, ret)
			flag++
		}
		currentNodeList = nextNodeList
	}

	return retList
}
