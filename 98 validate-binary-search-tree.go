package leetcode

import . "github.com/skyismine2010/leetcode/MyStruct"

func max(a []int) int {
	val := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > val {
			val = a[i]
		}
	}
	return val

}

func min(a []int) int {
	val := a[0]
	for i := 0; i < len(a); i++ {
		if a[i] < val {
			val = a[i]
		}
	}
	return val
}

func judgeBST(root *TreeNode) (bool, int, int) {
	var min, max int
	if root == nil {
		return true, 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return true, root.Val, root.Val
	}

	if root.Left != nil {
		ret, Lmax, Lmin := judgeBST(root.Left)
		if ret != true {
			return false, 0, 0
		}
		if Lmax >= root.Val {
			return false, 0, 0
		}

		min = Lmin
	}

	if root.Right != nil {
		ret, RMax, Rmin := judgeBST(root.Right)
		if ret != true {
			return false, 0, 0
		}
		if Rmin <= root.Val {
			return false, 0, 0
		}
		max = RMax
	}
	return true, max, min

}

func isValidBST(root *TreeNode) bool {
	ret, _, _ := judgeBST(root)
	return ret
}

func buildTree(list []int) *TreeNode {
	var nodeList []TreeNode
	for i := 0; i < len(list); i++ {
		nodeList = append(nodeList, TreeNode{list[i], nil, nil})
	}
	for i := 0; i < len(list)/2; i++ {
		nodeList[i].Left = &nodeList[i*2+1]
		nodeList[i].Right = &nodeList[i*2+2]
	}
	return &nodeList[0]

}
