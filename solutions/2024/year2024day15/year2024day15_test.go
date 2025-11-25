package year2024day15

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day15Part1(t *testing.T) {
	test.Assert(t, "1294459", part1, input)
}

func TestYear2024Day15Part2(t *testing.T) {
	test.Assert(t, "1319212", part2, input)
}
