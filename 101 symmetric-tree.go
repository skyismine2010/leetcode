package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil && root.Right != nil {
		return false
	}

	if root.Left != nil && root.Right == nil {
		return false
	}

	if root.Left.Val == root.Right.Val {
		LRet := isSymmetric(root.Left)
		if LRet != true {
			return false
		}
		RRet := isSymmetric(root.Right)
		if RRet != true {
			return false
		}
		return true
	}
	return false
}
