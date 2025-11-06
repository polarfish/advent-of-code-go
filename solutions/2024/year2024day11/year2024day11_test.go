package year2024day11

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day11Part1(t *testing.T) {
	test.Assert(t, "217812", part1, input)
}

func TestYear2024Day11Part2(t *testing.T) {
	test.Assert(t, "259112729857522", part2, input)
}
