package puzzles

import "testing"

func TestYear2015day02Part1(t *testing.T) {
	runTests(t, year2015Day02Part1, map[string]testCase{
		"input": {year2015Day02Input, "1606483"},
	})
}

func TestYear2015day02Part2(t *testing.T) {
	runTests(t, year2015Day02Part2, map[string]testCase{
		"input": {year2015Day02Input, "3842356"},
	})
}
