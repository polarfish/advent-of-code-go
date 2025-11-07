package year2024day12

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day12Part1(t *testing.T) {
	test.Assert(t, "1467094", part1, input)
}

func TestYear2024Day12Part2(t *testing.T) {
	test.Assert(t, "881182", part2, input)
}
