package year2025day05

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestYear2025Day05Part1Sample(t *testing.T) {
	test.Assert(t, "3", part1, sample)
}

func TestYear2025Day05Part1(t *testing.T) {
	test.Assert(t, "638", part1, input)
}

func TestYear2025Day05Part2Sample(t *testing.T) {
	test.Assert(t, "14", part2, sample)
}

func TestYear2025Day05Part2(t *testing.T) {
	test.Assert(t, "352946349407338", part2, input)
}
