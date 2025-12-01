package year2024day17

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day17Part1(t *testing.T) {
	test.Assert(t, "7,3,5,7,5,7,4,3,0", part1, input)
}

func TestYear2024Day17Part2(t *testing.T) {
	test.Assert(t, "105734774294938", part2, input)
}
