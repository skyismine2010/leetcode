package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func inorderIter(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if inorderIter(p.Left, q.Left) != true {
		return false
	}
	if inorderIter(p.Right, q.Right) != true {
		return false
	}

	return true

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return inorderIter(p, q)
}
