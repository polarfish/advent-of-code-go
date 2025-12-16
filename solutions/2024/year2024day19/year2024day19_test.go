package year2024day19

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

func TestYear2024Day19Part1(t *testing.T) {
	test.Assert(t, "336", part1, input)
}

func BenchmarkYear2024Day19Part1(b *testing.B) {
	for b.Loop() {
		part1(input)
	}
}

func TestYear2024Day19Part2(t *testing.T) {
	test.Assert(t, "758890600222015", part2, input)
}
