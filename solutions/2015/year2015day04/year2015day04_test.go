package year2015day04

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2015day04Part1(t *testing.T) {
	test.Assert(t, "282749", part1(input))
}

func TestYear2015day04Part2(t *testing.T) {
	test.Assert(t, "9962624", part2(input))
}
