package year2024day03

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day03Part1(t *testing.T) {
	test.Assert(t, "174960292", part1(input))
}

func TestYear2024Day03Part2(t *testing.T) {
	test.Assert(t, "56275602", part2(input))
}
