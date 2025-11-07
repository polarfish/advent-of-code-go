package year2024day13

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day13Part1(t *testing.T) {
	test.Assert(t, "29517", part1, input)
}

func TestYear2024Day13Part2(t *testing.T) {
	test.Assert(t, "103570327981381", part2, input)
}
