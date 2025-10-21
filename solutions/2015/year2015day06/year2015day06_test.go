package year2015day06

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2015Day06Part1(t *testing.T) {
	test.Assert(t, "400410", part1, input)
}

func TestYear2015Day06Part2(t *testing.T) {
	test.Assert(t, "15343601", part2, input)
}
