package year2024day02

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day02Part1(t *testing.T) {
	test.Assert(t, "334", part1, input)
}

func TestYear2024Day02Part2(t *testing.T) {
	test.Assert(t, "400", part2, input)
}
