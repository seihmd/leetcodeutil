package leetcodeutil

import (
	"regexp"
	"strconv"
	"strings"
)

type Interval struct {
	Start int
	End   int
}

var intervalPat = `\[[0-9]+,[0-9]+\]`
var intervalsRegexp = regexp.MustCompile(`^\[` + intervalPat + `(,` + intervalPat + `)*?\]$`)

// Intervals generates []Interval
func Intervals(input string) []Interval {
	if !intervalsRegexp.MatchString(input) {
		panic("invalid input")
	}
	s := strings.TrimPrefix(input, "[[")
	s = strings.TrimSuffix(s, "]]")

	intervalStrs := strings.Split(s, "],[")
	intervals := make([]Interval, len(intervalStrs))

	for i, intervalStr := range intervalStrs {
		numStrPair := strings.Split(intervalStr, ",")
		start, _ := strconv.Atoi(numStrPair[0])
		end, _ := strconv.Atoi(numStrPair[1])
		if start > end {
			panic("start must be lte end")
		}
		intervals[i] = Interval{
			Start: start,
			End:   end,
		}
	}
	return intervals
}
