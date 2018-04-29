package test

import (
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestBinaryTree(t *testing.T) {
	tn := leetcodeutil.BinaryTree("[-5,4,8,-11,null,13,-4,7,2,null,null,1]")

	tests := [][]int{
		[]int{tn.Val, -5},
		[]int{tn.Left.Val, 4},
		[]int{tn.Right.Val, 8},
		[]int{tn.Left.Left.Val, -11},
		[]int{tn.Right.Left.Val, 13},
		[]int{tn.Right.Right.Val, -4},
		[]int{tn.Left.Left.Left.Val, 7},
		[]int{tn.Left.Left.Right.Val, 2},
		[]int{tn.Right.Left.Left.Val, 1},
	}

	for _, test := range tests {
		if test[0] != test[1] {
			t.Errorf("TreeNode Val expect: %d, actual: %d", test[1], test[0])
		}
	}
	if tn.Left.Right != nil || tn.Right.Left.Right != nil {
		t.Error("node should be nil")
	}
}

func TestEmptyBinaryTree(t *testing.T) {
	tn := leetcodeutil.BinaryTree("[]")
	if tn != nil {
		t.Error("TreeNode should be nil")
	}

	tn = leetcodeutil.BinaryTree("[null]")
	if tn != nil {
		t.Error("*TreeNode should be nil")
	}
}

func TestInvalidBinaryTreeInput(t *testing.T) {
	tests := []string{"", "[", "]", "[[]", "[]]", "[null,1]", "[1,null,1,1]"}
	for _, test := range tests {
		failIfBinaryTreeNoPanic(test, t)
	}
}

func TestString(t *testing.T) {
	type testdata struct {
		input  string
		expect string
	}
	tests := []testdata{
		testdata{"[]", "[]"},
		testdata{"[null]", "[]"},
		testdata{"[1]", "[1]"},
		testdata{"[1,2,3]", "[1,2,3]"},
		testdata{"[1,2,3,null]", "[1,2,3]"},
		testdata{"[-1,2,-3]", "[-1,2,-3]"},
		testdata{"[5,4,8,11,null,13,4,7,2,null,null,1]", "[5,4,8,11,null,13,4,7,2,null,null,1]"},
	}
	for _, test := range tests {
		actual := leetcodeutil.BinaryTree(test.input).String()
		if actual != test.expect {
			t.Errorf("String expect: %s, actual: %s", test.expect, actual)
		}
	}
}

func failIfBinaryTreeNoPanic(s string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	leetcodeutil.BinaryTree(s)
}
