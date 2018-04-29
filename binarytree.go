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

func (root *TreeNode) String() string {
	if root == nil {
		return "[]"
	}
	vals := []string{}
	walk(root, 0, &vals)
	s := strings.Join(vals, ",")
	s = strings.TrimRight(s, ",")
	for {
		if strings.Contains(s, ",,") {
			s = strings.Replace(s, ",,", ",null,", -1)
		} else {
			break
		}
	}
	return "[" + s + "]"
}

func walk(n *TreeNode, i int, vals *[]string) {
	if n == nil {
		return
	}
	li := i*2 + 1
	ri := i*2 + 2
	if len(*vals)-1 < ri {
		emp := make([]string, ri-len(*vals))
		*vals = append(*vals, emp...)
	}
	(*vals)[i] = strconv.Itoa(n.Val)
	walk(n.Left, li, vals)
	walk(n.Right, ri, vals)
}

// Print outputs *TreeNode.String()
func (t *TreeNode) Print() {
	fmt.Println(t)
}
