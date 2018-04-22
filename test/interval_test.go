package test

import (
	"fmt"
	"testing"

	"github.com/seihmd/leetcodeutil"
)

func TestIntervals(t *testing.T) {
	type testdata struct {
		input        string
		sprintExpect string
	}
	tests := []testdata{
		{"[[1,2]]", "[{1 2}]"},
		{"[[2,2]]", "[{2 2}]"},
		{"[[12,34]]", "[{12 34}]"},
		{"[[1,2],[3,4]]", "[{1 2} {3 4}]"},
		{"[[12,34],[56,78]]", "[{12 34} {56 78}]"},
		{"[[11,12],[21,22],[31,32]]", "[{11 12} {21 22} {31 32}]"},
	}

	for _, test := range tests {
		m := leetcodeutil.Intervals(test.input)
		actual := fmt.Sprint(m)
		if test.sprintExpect != actual {
			t.Errorf("Intervals expect %s, actual %s", test.sprintExpect, actual)
		}
	}
}

func TestInvalidIntervalsInput(t *testing.T) {
	tests := []string{"", "[", "]", "[[]", "[]]", "[1]", "[[1]]", "[1,2]", "[[1, 2]]", "[[1,2],[1]]"}
	for _, test := range tests {
		failIfIntervalsNoPanic(test, t)
	}
}

func TestInvalidIntervals(t *testing.T) {
	tests := []string{
		"[[2,1]]",
	}
	for _, test := range tests {
		failIfIntervalsNoPanic(test, t)
	}
}

func failIfIntervalsNoPanic(s string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	leetcodeutil.Intervals(s)
}
