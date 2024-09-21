package binarytree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

// Recursive
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

// UnRecursive
func PreOrderUnRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	pop := func() *TreeNode {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return tmp
	}

	var res []int
	for len(stack) != 0 {
		root = pop()
		res = append(res, root.val)
		if root.right != nil {
			stack = append(stack, root.right)
		}
		if root.left != nil {
			stack = append(stack, root.left)
		}
	}
	return res
}

func InOrderUnRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	pop := func() *TreeNode {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return tmp
	}

	var res []int
	for len(stack) != 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			root = pop()
			res = append(res, root.val)
			root = root.right
		}
	}
	return res
}

func PosOrderUnRecur(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	pop := func() *TreeNode {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return tmp
	}

	var res []int
	for len(stack) != 0 {
		root = pop()
		res = append(res, root.val)
		if root.left != nil {
			stack = append(stack, root.left)
		}
		if root.right != nil {
			stack = append(stack, root.right)
		}
	}

	for l, r := 0, len(res)-1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

func LevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	pop := func() *TreeNode {
		tmp := queue[0]
		queue = queue[1:]
		return tmp
	}

	var res []int
	for len(queue) != 0 {
		root = pop()
		res = append(res, root.val)
		if root.left != nil {
			queue = append(queue, root.left)
		}
		if root.right != nil {
			queue = append(queue, root.right)
		}
	}
	return res
}

func LevelOrderWithLevelNo(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	bfs(0, root, &res)
	return res
}

func bfs(level int, root *TreeNode, res *[][]int) {
	if root == nil {
		return
	}

	if len(*res) < level+1 {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.val)
	if root.left != nil {
		bfs(level+1, root.left, res)
	}
	if root.right != nil {
		bfs(level+1, root.right, res)
	}
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.val != q.val {
		return false
	}
	return IsSameTree(p.left, q.left) && IsSameTree(p.right, q.right)
}

func GenBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return genBST(nums, 0, len(nums)-1)
}

func genBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := &TreeNode{val: nums[mid]}
	root.left = genBST(nums, start, mid-1)
	root.right = genBST(nums, mid+1, end)
	return root
}
