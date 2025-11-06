package year2024day10

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day10Part1(t *testing.T) {
	test.Assert(t, "607", part1, input)
}

func TestYear2024Day10Part2(t *testing.T) {
	test.Assert(t, "1384", part2, input)
}
