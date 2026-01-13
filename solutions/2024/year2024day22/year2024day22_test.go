package year2024day22

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day22Part1(t *testing.T) {
	test.Assert(t, "14082561342", part1, input)
}

func TestYear2024Day22Part2(t *testing.T) {
	test.Assert(t, "1568", part2, input)
}
