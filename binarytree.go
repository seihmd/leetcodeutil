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

const treeNodePat = "(-?[0-9]+|null)"

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

	dummyRoot := &TreeNode{}
	queue := []*TreeNode{dummyRoot}
	for i, val := range vals {
		node := queue[0]
		if i%2 == 0 {
			queue = queue[1:]
		}
		if val == "null" {
			continue
		}

		n, _ := strconv.Atoi(val)
		child := &TreeNode{Val: n}
		if i%2 != 0 {
			node.Left = child
		} else {
			node.Right = child
		}
		queue = append(queue, child)
	}
	return dummyRoot.Right
}

func (t *TreeNode) String() string {
	if t == nil {
		return "[]"
	}
	vals := []string{}
	dummyRoot := &TreeNode{Right: t}
	queue := []*TreeNode{dummyRoot}
	notNilCnt := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			notNilCnt--
		}

		vals = append(vals, node.toStr())
		if node == nil {
			continue
		}
		if node.Left != nil {
			notNilCnt++
		}
		if node.Right != nil {
			notNilCnt++
		}
		if notNilCnt == 0 {
			break
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	return "[" + strings.Join(vals[2:], ",") + "]"
}

func (t *TreeNode) toStr() string {
	if t == nil {
		return "null"
	}
	return strconv.Itoa(t.Val)
}

// Print outputs *TreeNode.String()
func (t *TreeNode) Print() {
	fmt.Println(t)
}
