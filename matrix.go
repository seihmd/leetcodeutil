package leetcodeutil

import (
	"regexp"
	"strconv"
	"strings"
)

var matrixRowPat = `\[([0-9+]+((,[0-9]+)+)?)?\]`
var matrixInputPat = regexp.MustCompile(`^\[(` + matrixRowPat + `)?(,` + matrixRowPat + `)*?\]$`)

// Matrix generates [][]int from string like "[[1,2,3],[4,5,6]]"
func Matrix(input string) [][]int {
	if !matrixInputPat.MatchString(input) {
		panic("invalid input")
	}

	input = remBrackets(input)
	sliceStrs := strings.Split(input, "],[")

	if len(sliceStrs[0]) == 0 {
		return [][]int{}
	}

	matrix := make([][]int, len(sliceStrs))
	for i, sliceStr := range sliceStrs {
		sliceStr = remBrackets(sliceStr)
		matrix[i] = toInts(sliceStr)
		if i > 0 && len(matrix[i-1]) != len(matrix[i]) {
			panic("rows must have same length")
		}
	}

	return matrix
}

func remBrackets(s string) string {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	return s
}

func toInts(s string) []int {
	if len(s) == 0 {
		return []int{}
	}
	nums := strings.Split(s, ",")
	ints := make([]int, len(nums))
	for i, num := range nums {
		n, _ := strconv.Atoi(num)
		ints[i] = n
	}
	return ints
}
