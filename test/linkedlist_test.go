package test

import (
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestEmptyLinkedList(t *testing.T) {
	l := leetcodeutil.LinkedList(nil)
	if l != nil {
		t.Error("LinkedList should be nil")
	}
}

func TestLinkedList(t *testing.T) {
	l := leetcodeutil.LinkedList([]int{1, 2, 3})

	type testdata struct {
		listnode *leetcodeutil.ListNode
		expect   int
	}
	tests := []testdata{
		{l, 1},
		{l.Next, 2},
		{l.Next.Next, 3},
	}

	for i, test := range tests {
		if test.listnode.Val != test.expect {
			t.Errorf("ListNode: %d, Val expects %d, actual %d ", i, test.expect, test.listnode.Val)
		}
	}
}

func TestEmptyLinkedListString(t *testing.T) {
	l := leetcodeutil.LinkedList(nil)
	if l.String() != "<nil>" {
		t.Fail()
	}
}

func TestListNodeString(t *testing.T) {
	l := leetcodeutil.LinkedList([]int{1, 2, 3})
	if l.String() != "1 2 3" {
		t.Errorf("LinkedList.String expects '1 2 3', actual '%s'", l.String())
	}
}
