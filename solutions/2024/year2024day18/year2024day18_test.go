package year2024day18

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day18Part1(t *testing.T) {
	test.Assert(t, "322", part1, input)
}

func TestYear2024Day18Part2(t *testing.T) {
	test.Assert(t, "60,21", part2, input)
}
