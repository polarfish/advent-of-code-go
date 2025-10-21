package year2015day01

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2015day01Part1(t *testing.T) {
	test.Assert(t, "280", part1, input)
}

func TestYear2015day01Part2(t *testing.T) {
	test.Assert(t, "1797", part2, input)
}
