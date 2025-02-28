package puzzles

import "testing"

func TestYear2015day05Part1(t *testing.T) {
	runTests(t, year2015Day05Part1, map[string]testCase{
		"input": {year2015Day05Input, "255"},
	})
}

func TestYear2015day05Part2(t *testing.T) {
	runTests(t, year2015Day05Part2, map[string]testCase{
		"input": {year2015Day05Input, "55"},
	})
}
