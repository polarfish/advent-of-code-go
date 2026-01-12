package year2024day21

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `
029A
980A
179A
456A
379A
`

func TestYear2024Day21Part1(t *testing.T) {
	test.Assert(t, "126384", part1, sample)
	test.Assert(t, "156714", part1, input)
}

func TestYear2024Day21Part2(t *testing.T) {
	test.Assert(t, "154115708116294", part2, sample)
	test.Assert(t, "191139369248202", part2, input)
}
