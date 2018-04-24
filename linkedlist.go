package leetcodeutil

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList generates *ListNode
func LinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	root := &ListNode{
		Val: vals[0],
	}
	current := root
	for i := 1; i < len(vals); i++ {
		next := &ListNode{
			Val: vals[i],
		}
		current.Next = next
		current = next
	}
	return root
}

func (ln *ListNode) String() (r string) {
	if ln == nil {
		return fmt.Sprint(nil)
	}
	c := ln
	for c != nil {
		r += strconv.Itoa(c.Val) + " "
		c = c.Next
	}
	r = strings.TrimSuffix(r, " ")
	return
}

// Print outputs *ListNode.String()
func (ln *ListNode) Print() {
	fmt.Println(ln)
}
