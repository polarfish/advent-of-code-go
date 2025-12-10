package year2025day10

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
`

func TestYear2025Day10Part1Sample(t *testing.T) {
	test.Assert(t, "7", part1, sample)
}

func TestYear2025Day10Part1(t *testing.T) {
	test.Assert(t, "481", part1, input)
}

func TestYear2025Day10Part2Sample(t *testing.T) {
	test.Assert(t, "33", part2, sample)
}

func TestYear2025Day10Part2(t *testing.T) {
	test.Assert(t, "0", part2, input)
}
