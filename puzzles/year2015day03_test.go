package puzzles

import "testing"

func TestYear2015day03Part1(t *testing.T) {
	runTests(t, year2015Day03Part1, map[string]testCase{
		"input": {year2015Day03Input, "2572"},
	})
}

func TestYear2015day03Part2(t *testing.T) {
	runTests(t, year2015Day03Part2, map[string]testCase{
		"input": {year2015Day03Input, "2631"},
	})
}
