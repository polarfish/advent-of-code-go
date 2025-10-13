package year2015day05

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2015day05Part1(t *testing.T) {
	test.Assert(t, "255", part1(input))
}

func TestYear2015day05Part2(t *testing.T) {
	test.Assert(t, "55", part2(input))
}
