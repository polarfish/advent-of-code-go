package year2025day03

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
987654321111111
811111111111119
234234234234278
818181911112111
`

func TestYear2025Day03Part1Sample(t *testing.T) {
	test.Assert(t, "357", part1, sample)
}

func TestYear2025Day03Part1(t *testing.T) {
	test.Assert(t, "17229", part1, input)
}

func TestYear2025Day03Part2Sample(t *testing.T) {
	test.Assert(t, "3121910778619", part2, sample)
}

func TestYear2025Day03Part2(t *testing.T) {
	test.Assert(t, "170520923035051", part2, input)
}
