package year2025day12

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
`

func TestYear2025Day12Part1Sample(t *testing.T) {
	test.Assert(t, "1", part1, sample) // it should be 2
}

func TestYear2025Day12Part1(t *testing.T) {
	test.Assert(t, "567", part1, input)
}

func TestYear2025Day12Part2(t *testing.T) {
	test.Assert(t, "0", part2, input)
}
