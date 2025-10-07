package year2024day04

import (
	"testing"

	"github.com/polarfish/advent-of-code-go/puzzles/utils"
)

func TestYear2024Day04Part1(t *testing.T) {
	utils.Test(t, "2336", part1(input))
}

func TestYear2024Day04Part2(t *testing.T) {
	utils.Test(t, "1831", part2(input))
}
