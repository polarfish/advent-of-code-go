package puzzles

import "testing"

func TestYear2015day01Part1(t *testing.T) {
	runTests(t, year2015Day01Part1, map[string]testCase{
		"input": {year2015Day01Input, "280"},
	})
}

func TestYear2015day01Part2(t *testing.T) {
	runTests(t, year2015Day01Part2, map[string]testCase{
		"input": {year2015Day01Input, "1797"},
	})
}
