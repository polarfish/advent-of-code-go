package year2024day14

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day14Part1(t *testing.T) {
	test.Assert(t, "230686500", part1, input)
}

func TestYear2024Day14Part2(t *testing.T) {
	test.Assert(t, "7672", part2, input)
}
