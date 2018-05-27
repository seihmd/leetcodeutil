package leetcodeutil

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const listNodePattern = `-?[0-9]+`

var linkedListInputPattern = regexp.MustCompile(`^\[(` + listNodePattern + `)?(,` + listNodePattern + `)*?\]$`)

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedList generates *ListNode
func LinkedList(input string) *ListNode {
	if !linkedListInputPattern.MatchString(input) {
		panic("invalid input")
	}

	input = strings.TrimPrefix(input, "[")
	input = strings.TrimSuffix(input, "]")
	if len(input) == 0 {
		return nil
	}
	vals := strings.Split(input, ",")

	dummyHead := &ListNode{}
	current := dummyHead
	for _, val := range vals {
		n, _ := strconv.Atoi(val)
		next := &ListNode{
			Val: n,
		}
		current.Next = next
		current = next
	}
	return dummyHead.Next
}

func (ln *ListNode) String() (r string) {
	if ln == nil {
		return "[]"
	}
	c := ln
	for c != nil {
		r += strconv.Itoa(c.Val) + " "
		c = c.Next
	}
	r = strings.TrimSuffix(r, " ")
	r = "[" + r + "]"
	return
}

// Print outputs *ListNode.String()
func (ln *ListNode) Print() {
	fmt.Println(ln)
}
