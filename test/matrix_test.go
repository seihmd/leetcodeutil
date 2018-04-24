package test

import (
	"fmt"
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestMatrix(t *testing.T) {
	type testdata struct {
		input        string
		sprintExpect string
	}
	tests := []testdata{
		{"[[1]]", "[[1]]"},
		{"[[123]]", "[[123]]"},
		{"[[1,2,3]]", "[[1 2 3]]"},
		{"[[12,34,56]]", "[[12 34 56]]"},
		{"[[1,2,3],[4,5,6],[7,8,9]]", "[[1 2 3] [4 5 6] [7 8 9]]"},
		{"[[1,2],[3,4],[5,6]]", "[[1 2] [3 4] [5 6]]"},
		{"[[1,2,3],[4,5,6]]", "[[1 2 3] [4 5 6]]"},
		{"[[1,2,3],[45,67,89]]", "[[1 2 3] [45 67 89]]"},
		{"[[12,34,56],[7,8,9]]", "[[12 34 56] [7 8 9]]"},
		{"[]", "[]"},
		{"[[],[],[]]", "[[] [] []]"},
	}

	for _, test := range tests {
		m := leetcodeutil.Matrix(test.input)
		actual := fmt.Sprint(m)
		if test.sprintExpect != actual {
			t.Errorf("Matrix expect %s, actual %s", test.sprintExpect, actual)
		}
	}
}

func TestInvalidInput(t *testing.T) {
	tests := []string{"", "[", "]", "[[]", "[]]", "[1]", "[1,2]", "[[1,2, 3]]", "[[1], [2]]"}
	for _, test := range tests {
		failIfMatrixNoPanic(test, t)
	}
}

func TestInvalidMatrix(t *testing.T) {
	tests := []string{
		"[[1,2][3]]",
		"[[1],[2,3]]",
		"[[],[1]]",
	}
	for _, test := range tests {
		failIfMatrixNoPanic(test, t)
	}
}

func failIfMatrixNoPanic(s string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	leetcodeutil.Matrix(s)
}
