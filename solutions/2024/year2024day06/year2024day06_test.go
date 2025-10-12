package year2024day06

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day06Part1(t *testing.T) {
	test.Assert(t, "5531", part1(input))
}

func TestYear2024Day06Part2(t *testing.T) {
	test.Assert(t, "2165", part2(input))
}
