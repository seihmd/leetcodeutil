package test

import (
	"testing"

	. "github.com/seihmd/leetcodeutil"
)

func TestBinaryTree(t *testing.T) {
	type testdata struct {
		tn      *TreeNode
		compare *TreeNode
	}

	tests := []testdata{
		testdata{BinaryTree("[]"), nil},
		testdata{BinaryTree("[null]"), nil},
		testdata{BinaryTree("[1]"), &TreeNode{Val: 1, Left: nil, Right: nil}},
		testdata{BinaryTree("[1,null,2,null,3]"), &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3, Left: nil, Right: nil}}}},
		testdata{BinaryTree("[1,2,null,3]"), &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}, Right: nil}},
		testdata{BinaryTree("[-5,4,8,-11,null,13,-4,7,2,null,null,1]"), &TreeNode{Val: -5,
			Left: &TreeNode{Val: 4,
				Left: &TreeNode{Val: -11,
					Left:  &TreeNode{Val: 7, Left: nil, Right: nil},
					Right: &TreeNode{Val: 2, Left: nil, Right: nil},
				},
				Right: nil,
			},
			Right: &TreeNode{Val: 8,
				Left: &TreeNode{Val: 13, Left: nil, Right: nil},
				Right: &TreeNode{Val: -4,
					Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
					Right: nil,
				},
			},
		}},
	}

	for i, test := range tests {
		if !isSame(test.tn, test.compare) {
			t.Errorf("%dth test failed", i)
		}
	}
}

func isSame(tn *TreeNode, compare *TreeNode) bool {
	if tn == nil || compare == nil {
		return tn == nil && compare == nil
	}
	if tn.Val != compare.Val {
		return false
	}
	return isSame(tn.Left, compare.Left) &&
		isSame(tn.Right, compare.Right)
}

func TestEmptyBinaryTree(t *testing.T) {
	tn := BinaryTree("[]")
	if tn != nil {
		t.Error("TreeNode should be nil")
	}

	tn = BinaryTree("[null]")
	if tn != nil {
		t.Error("*TreeNode should be nil")
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
		testdata{"[1,2,null,3]", "[1,2,null,3]"},
		testdata{"[1,null,2,null,3]", "[1,null,2,null,3]"},
		testdata{"[1]", "[1]"},
		testdata{"[-1]", "[-1]"},
		testdata{"[-1,null]", "[-1]"},
		testdata{"[1,2,3]", "[1,2,3]"},
		testdata{"[1,2,3,null]", "[1,2,3]"},
		testdata{"[-1,2,-3]", "[-1,2,-3]"},
		testdata{"[5,4,8,11,null,13,4,7,2,null,null,1]", "[5,4,8,11,null,13,4,7,2,null,null,1]"},
	}
	for _, test := range tests {
		actual := BinaryTree(test.input).String()
		if actual != test.expect {
			t.Errorf("\nString expect: %s,\n actual: %s", test.expect, actual)
		}
	}
}

func TestInvalidBinaryTreeInput(t *testing.T) {
	tests := []string{"", "[", "]", "[[]", "[]]", "[null,1]", "[1,null,null,null]"}
	for _, test := range tests {
		failIfBinaryTreeNoPanic(test, t)
	}
}

func failIfBinaryTreeNoPanic(s string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	BinaryTree(s)
}
