package leetcodeutil

import (
	"strconv"
	"strings"
)

func IntSlice(input string) []int {
	if !matrixRowRegexp.MatchString(input) {
		panic("invalid input")
	}

	input = remBrackets(input)
	numStrs := strings.Split(input, ",")

	if numStrs[0] == "" {
		return []int{}
	}

	intSlice := make([]int, len(numStrs))
	for i, numStr := range numStrs {
		intSlice[i], _ = strconv.Atoi(numStr)
	}
	return intSlice
}
