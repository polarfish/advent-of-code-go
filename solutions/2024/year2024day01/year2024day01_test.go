package year2024day01

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestYear2024Day01Part1Sample(t *testing.T) {
	test.Assert(t, "11", part1, sample)
}

func TestYear2024Day01Part1(t *testing.T) {
	test.Assert(t, "2430334", part1, input)
}

func TestYear2024Day01Part2Sample(t *testing.T) {
	test.Assert(t, "31", part2, sample)
}

func TestYear2024Day01Part2(t *testing.T) {
	test.Assert(t, "28786472", part2, input)
}
