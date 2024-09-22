package binarytree

import (
	"github.com/bytedance/sonic"
	"testing"
)

func initBinaryTree(t *testing.T) *TreeNode {
	t.Helper()

	root := NewTreeNode(1)
	root.left = NewTreeNode(2)
	root.right = NewTreeNode(3)
	root.left.left = NewTreeNode(4)
	root.left.right = NewTreeNode(5)
	root.right.left = NewTreeNode(6)
	root.right.right = NewTreeNode(7)
	return root
}

func TestPreOrderRecur(t *testing.T) {
	res := PreOrderRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestInOrderRecur(t *testing.T) {
	res := InOrderRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestPosOrderRecur(t *testing.T) {
	res := PosOrderRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestPreOrderUnRecur(t *testing.T) {
	res := PreOrderUnRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestInOrderUnRecur(t *testing.T) {
	res := InOrderUnRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestPosOrderUnRecur(t *testing.T) {
	res := PosOrderUnRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestLevelOrder(t *testing.T) {
	res := LevelOrder(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestLevelOrderWithLevelNo(t *testing.T) {
	res := LevelOrderWithLevelNoRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %s", bs)
}

func TestLevelOrderWithLevelNoUnRecur(t *testing.T) {
	res := LevelOrderWithLevelNoUnRecur(initBinaryTree(t))
	bs, _ := sonic.MarshalString(res)
	t.Logf("Res: %+v", bs)
}

func Test_IsSameTree(t *testing.T) {
	p := initBinaryTree(t)
	q := initBinaryTree(t)

	b := IsSameTree(p, q)
	t.Logf("Res: %+v", b)
}

func TestGenerateBSTree(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	root := GenBST(nums)
	t.Logf("Res: %+v", root)
}
