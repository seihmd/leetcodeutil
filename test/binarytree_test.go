package test

import (
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestBinaryTree(t *testing.T) {
	tn := leetcodeutil.BinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8})

	tests := [][]int{
		[]int{tn.Val, 1},
		[]int{tn.Left.Val, 2},
		[]int{tn.Right.Val, 3},
		[]int{tn.Left.Left.Val, 4},
		[]int{tn.Left.Right.Val, 5},
		[]int{tn.Right.Left.Val, 6},
		[]int{tn.Right.Right.Val, 7},
		[]int{tn.Left.Left.Left.Val, 8},
	}

	for _, test := range tests {
		if test[0] != test[1] {
			t.Errorf("TreeNode Val expect: %d, actual: %d", test[1], test[0])
		}
	}
}

func TestEmptyBinaryTree(t *testing.T) {
	tn := leetcodeutil.BinaryTree([]int{})
	if tn != nil {
		t.Error("TreeNode should be nil")
	}
}
