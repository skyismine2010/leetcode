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
func levelOrder(root *TreeNode) [][]int {
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
			retList = append(retList, ret)
		}
		currentNodeList = nextNodeList
	}

	return retList
}
