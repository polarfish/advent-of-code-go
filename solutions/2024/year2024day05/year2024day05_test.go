package year2024day05

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day05Part1(t *testing.T) {
	test.Assert(t, "5713", part1(input))
}

func TestYear2024Day05Part2(t *testing.T) {
	test.Assert(t, "5180", part2(input))
}
