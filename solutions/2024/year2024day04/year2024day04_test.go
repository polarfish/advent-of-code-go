package year2024day04

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day04Part1(t *testing.T) {
	test.Assert(t, "2336", part1, input)
}

func TestYear2024Day04Part2(t *testing.T) {
	test.Assert(t, "1831", part2, input)
}
