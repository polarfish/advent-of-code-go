package puzzles

import "testing"

func TestYear2015Day06Part1(t *testing.T) {
	runTests(t, year2015Day06Part1, map[string]testCase{
		"input": {year2015Day06Input, "400410"},
	})
}

func TestYear2015Day06Part2(t *testing.T) {
	runTests(t, year2015Day06Part2, map[string]testCase{
		"input": {year2015Day06Input, "15343601"},
	})
}
