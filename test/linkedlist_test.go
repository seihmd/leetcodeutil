package test

import (
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestLinkedList(t *testing.T) {
	l := leetcodeutil.LinkedList("[1,2,3]")

	type testdata struct {
		listnode *leetcodeutil.ListNode
		expect   int
	}
	tests := []testdata{{l, 1}, {l.Next, 2}, {l.Next.Next, 3}}

	for i, test := range tests {
		if test.listnode.Val != test.expect {
			t.Errorf("ListNode: %d, Val expects %d, actual %d ", i, test.expect, test.listnode.Val)
		}
	}
}

func TestEmptyLinkedList(t *testing.T) {
	l := leetcodeutil.LinkedList("[]")
	if l != nil {
		t.Error("LinkedList should be nil")
	}
}

func TestLinkedListString(t *testing.T) {
	type testdata struct {
		input  string
		expect string
	}
	tests := []testdata{
		testdata{"[]", "[]"},
		testdata{"[1]", "[1]"},
		testdata{"[1,2,3]", "[1 2 3]"},
		testdata{"[-1,-2,-3]", "[-1 -2 -3]"},
	}
	for _, test := range tests {
		l := leetcodeutil.LinkedList(test.input)
		if l.String() != test.expect {
			t.Errorf("LinkedList.String expects '%s', actual '%s'", test.expect, l.String())
		}
	}
}
