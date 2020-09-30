package test

import (
	"fmt"
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestIntSlice(t *testing.T) {
	type testdata struct {
		input        string
		sprintExpect string
	}
	tests := []testdata{
		{"[1]", "[1]"},
		{"[2,3]", "[2 3]"},
		{"[123]", "[123]"},
		{"[-1,-2,-3]", "[-1 -2 -3]"},
		{"[12,34,56]", "[12 34 56]"},
		{"[1,-2,3]", "[1 -2 3]"},
		{"[]", "[]"},
	}

	for _, test := range tests {
		m := leetcodeutil.IntSlice(test.input)
		actual := fmt.Sprint(m)
		if test.sprintExpect != actual {
			t.Errorf("IntSlice expect %s, actual %s", test.sprintExpect, actual)
		}
	}
}

func TestInvalidIntSliceInput(t *testing.T) {
	tests := []string{
		"", "[", "]",
		"[1, 2]",
		"[-]", "[-1,-]", "[-,-2]"}
	for _, test := range tests {
		failIfIntSliceNoPanic(test, t)
	}
}

func failIfIntSliceNoPanic(s string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	leetcodeutil.IntSlice(s)
}
