package leetcodeutil

import (
	"strings"
)

func IntSlices(input string) [][]int {
	if !matrixInputPat.MatchString(input) {
		panic("invalid input")
	}

	input = remBrackets(input)
	sliceStrs := strings.Split(input, "],[")

	if len(sliceStrs[0]) == 0 {
		return [][]int{}
	}

	intSlices := make([][]int, len(sliceStrs))
	for i, sliceStr := range sliceStrs {
		sliceStr = remBrackets(sliceStr)
		intSlices[i] = toInts(sliceStr)
	}
	return intSlices
}
