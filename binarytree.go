package leetcodeutil

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Binarytree generates TreeNode from []int
func Binarytree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	n := &TreeNode{Val: vals[0]}
	m := make([]*TreeNode, len(vals))
	m[0] = n
	for i := 1; i < len(vals); i++ {
		t := &TreeNode{
			Val: vals[i],
		}
		m[i] = t
		if i%2 == 0 {
			m[(i-2)/2].Right = t
		} else {
			m[(i-1)/2].Left = t
		}
	}
	return n
}
