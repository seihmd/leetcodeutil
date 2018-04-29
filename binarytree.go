package leetcodeutil

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const treeNodePat = "([0-9]+|null)"

var binarytreeInputPat = regexp.MustCompile(`^\[(` + treeNodePat + `((,` + treeNodePat + `)+)?)?\]$`)

// BinaryTree generates *TreeNode
func BinaryTree(input string) *TreeNode {
	if !binarytreeInputPat.MatchString(input) {
		panic("invalid input")
	}
	input = strings.TrimLeft(input, "[")
	input = strings.TrimRight(input, "]")
	if len(input) == 0 {
		return nil
	}
	vals := strings.Split(input, ",")

	m := make([]*TreeNode, len(vals))
	for i, val := range vals {
		var tn *TreeNode
		if val != "null" {
			intVal, _ := strconv.Atoi(vals[i])
			tn = &TreeNode{
				Val: intVal,
			}
		}
		m[i] = tn
		if i == 0 {
			continue
		} else if tn == nil {
			continue
		} else if i%2 == 0 {
			parent := m[(i-2)/2]
			if parent == nil {
				panic(fmt.Sprintf("parent of %dth node is nil", i))
			}
			parent.Right = tn
		} else {
			parent := m[(i-1)/2]
			if parent == nil {
				panic(fmt.Sprintf("parent of %dth node is nil", i))
			}
			parent.Left = tn
		}
	}
	return m[0]
}

func (t *TreeNode) String() string {
	nilVal := fmt.Sprint(nil)
	if t == nil {
		return nilVal
	}
	l, r := nilVal, nilVal
	if t.Left != nil {
		l = strconv.Itoa(t.Left.Val)
	}
	if t.Right != nil {
		r = strconv.Itoa(t.Right.Val)
	}
	return fmt.Sprintf("Val: %d, Left: %s, Right: %s", t.Val, l, r)
}

// Print outputs *TreeNode.String()
func (t *TreeNode) Print() {
	fmt.Println(t)
}
