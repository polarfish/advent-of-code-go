package puzzles

import "testing"

func TestYear2015Day07Part1(t *testing.T) {
	runTests(t, year2015Day07Part1, map[string]testCase{
		"input": {year2015Day07Input, "3176"},
	})
}

func TestYear2015Day07Part2(t *testing.T) {
	runTests(t, year2015Day07Part2, map[string]testCase{
		"input": {year2015Day07Input, "14710"},
	})
}
