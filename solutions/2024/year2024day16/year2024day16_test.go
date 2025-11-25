package year2024day16

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day16Part1(t *testing.T) {
	test.Assert(t, "89460", part1, input)
}

func TestYear2024Day16Part2(t *testing.T) {
	test.Assert(t, "504", part2, input)
}
