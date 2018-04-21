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
		{"[[1,2,3]]", "[[1 2 3]]"},
		{"[[1,2,3],[4,5,6],[7,8,9]]", "[[1 2 3] [4 5 6] [7 8 9]]"},
		{"[[1,2],[3,4],[5,6]]", "[[1 2] [3 4] [5 6]]"},
		{"[[1,2,3],[4,5,6]]", "[[1 2 3] [4 5 6]]"},
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
		failIfNoPanic(test, t)
	}
}

func TestInvalidMatrix(t *testing.T) {
	tests := []string{
		"[[1,2][3]]",
		"[[1],[2,3]]",
		"[[],[1]]",
	}
	for _, test := range tests {
		failIfNoPanic(test, t)
	}
}

func failIfNoPanic(s string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	leetcodeutil.Matrix(s)
}
