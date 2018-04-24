package test

import (
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestBinaryTree(t *testing.T) {
	tn := leetcodeutil.BinaryTree("[1,2,3,4,5,6,7,8]")

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

	tn = leetcodeutil.BinaryTree("[1,null,2]")
	if tn.Val != 1 || tn.Left != nil || tn.Right.Val != 2 {
		t.Fail()
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
		testdata{"[]", "<nil>"},
		testdata{"[1]", "Val: 1, Left: <nil>, Right: <nil>"},
		testdata{"[1,2]", "Val: 1, Left: 2, Right: <nil>"},
		testdata{"[1,2,3]", "Val: 1, Left: 2, Right: 3"},
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
