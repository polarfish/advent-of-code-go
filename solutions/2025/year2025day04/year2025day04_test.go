package year2025day04

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestYear2025Day04Part1Sample(t *testing.T) {
	test.Assert(t, "13", part1, sample)
}

func TestYear2025Day04Part1(t *testing.T) {
	test.Assert(t, "1464", part1, input)
}

func TestYear2025Day04Part2Sample(t *testing.T) {
	test.Assert(t, "43", part2, sample)
}

func TestYear2025Day04Part2(t *testing.T) {
	test.Assert(t, "8409", part2, input)
}
