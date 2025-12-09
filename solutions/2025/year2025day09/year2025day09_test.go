package year2025day09

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
`

func TestYear2025Day09Part1Sample(t *testing.T) {
	test.Assert(t, "50", part1, sample)
}

func TestYear2025Day09Part1(t *testing.T) {
	test.Assert(t, "4755429952", part1, input)
}

func TestYear2025Day09Part2Sample(t *testing.T) {
	test.Assert(t, "24", part2, sample)
}

func TestYear2025Day09Part2(t *testing.T) {
	test.Assert(t, "1429596008", part2, input)
}
