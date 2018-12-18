package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func inorderTraversal(root *TreeNode) []int {
	var ret []int

	if root.Left != nil {
		ret = append(ret, inorderTraversal(root.Left)...)
	}

	ret = append(ret, root.Val)

	if root.Right != nil {
		ret = append(ret, inorderTraversal(root.Right)...)
	}

	return ret
}
