package year2024day09

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/tools/test"
)

var sample = `2333133121414131402`

func TestYear2024Day09Part1(t *testing.T) {
	test.Assert(t, "1928", part1, sample)
	test.Assert(t, "6288599492129", part1, input)
}

func TestYear2024Day09Part2(t *testing.T) {
	test.Assert(t, "2858", part2, sample)
	test.Assert(t, "6321896265143", part2, input)
}
