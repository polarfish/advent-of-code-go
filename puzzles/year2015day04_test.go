package puzzles

import "testing"

func TestYear2015day04Part1(t *testing.T) {
	runTests(t, year2015Day04Part1, map[string]testCase{
		"input": {year2015Day04Input, "282749"},
	})
}

func TestYear2015day04Part2(t *testing.T) {
	runTests(t, year2015Day04Part2, map[string]testCase{
		"input": {year2015Day04Input, "9962624"},
	})
}
