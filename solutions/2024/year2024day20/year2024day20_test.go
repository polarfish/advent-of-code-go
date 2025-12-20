package year2024day20

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`

func TestYear2024Day20Part1Sample(t *testing.T) {
	test.Assert(t, "0", part1, sample) // 0, because there are not shortcuts more than 100 in the sample
}

func TestYear2024Day20Part1(t *testing.T) {
	test.Assert(t, "1317", part1, input)
}

func TestYear2024Day20Part2Sample(t *testing.T) {
	test.Assert(t, "0", part2, sample) // 0, because there are not shortcuts more than 100 in the sample
}

func TestYear2024Day20Part2(t *testing.T) {
	test.Assert(t, "982474", part2, input)
}
