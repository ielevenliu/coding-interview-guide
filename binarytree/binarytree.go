package binarytree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

func PreOrderRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	preOrderRecur(root, &res)
	return res
}

func preOrderRecur(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.val)
	preOrderRecur(root.left, res)
	preOrderRecur(root.right, res)
}

func InOrderRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	inOrderRecur(root, &res)
	return res
}

func inOrderRecur(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrderRecur(root.left, res)
	*res = append(*res, root.val)
	inOrderRecur(root.right, res)
}

func PosOrderRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	posOrderRecur(root, &res)
	return res
}

func posOrderRecur(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	posOrderRecur(root.left, res)
	posOrderRecur(root.right, res)
	*res = append(*res, root.val)
}
